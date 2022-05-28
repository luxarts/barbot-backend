package router

import (
	"barbot/internal/controller"
	"barbot/internal/defines"
	"barbot/internal/middleware"
	"barbot/internal/repository"
	"barbot/internal/service"
	"log"
	"os"
	"path"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	mdRepoFile := os.Getenv(defines.EnvDBFilePathMixedDrinks)
	if mdRepoFile == "" {
		log.Fatalf("%s empty", defines.EnvDBFilePathMixedDrinks)
	}
	drinksRepoFile := os.Getenv(defines.EnvDBFilePathDrinks)
	if drinksRepoFile == "" {
		log.Fatalf("%s empty", defines.EnvDBFilePathDrinks)
	}

	// Repositories
	mdRepo := repository.NewMixedDrinksRepository(path.Join(wd, mdRepoFile))
	drinksRepo := repository.NewDrinksRepository(path.Join(wd, drinksRepoFile))
	usersRepo := repository.NewUsersRepository()

	// Services
	mdSrv := service.NewMixedDrinksService(mdRepo)
	drinksSrv := service.NewDrinksService(drinksRepo)
	usersSrv := service.NewUsersService(usersRepo)
	sessionsSrv := service.NewSessionsService(usersRepo)

	// Controllers
	mdCtrl := controller.NewMixedDrinksController(mdSrv)
	drinksCtrl := controller.NewDrinksController(drinksSrv)
	usersCtrl := controller.NewUsersController(usersSrv)
	sessionsCtrl := controller.NewSessionsController(sessionsSrv)

	// Middlewares
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Header", "Access-Control-Request-Headers", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))
	mwAuth := middleware.NewAuthMiddleware()

	// Endpoints
	r.GET(defines.EndpointGetMixedDrinks, mdCtrl.GetAll)
	r.GET(defines.EndpointGetDrinks, drinksCtrl.GetAll)
	r.POST(defines.EndpointCreateUser, usersCtrl.Create)
	r.GET(defines.EndpointGetUserByUUID, mwAuth.Check, usersCtrl.GetByUUID)
	r.POST(defines.EndpointCreateSession, sessionsCtrl.Create)
	r.GET(defines.EndpointGetSession, mwAuth.Check, sessionsCtrl.Get)
}
