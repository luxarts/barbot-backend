package router

import (
	"barbot/internal/controller"
	"barbot/internal/defines"
	"barbot/internal/middleware"
	"barbot/internal/repository"
	"barbot/internal/service"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// Repositories
	mdRepo := repository.NewMixedDrinksRepository()
	drinksRepo := repository.NewDrinksRepository()
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
	r.Use(cors.Default())
	mwAuth := middleware.NewAuthMiddleware()

	// Endpoints
	r.GET(defines.EndpointGetMixedDrinks, mdCtrl.GetAll)
	r.GET(defines.EndpointGetDrinks, drinksCtrl.GetAll)
	r.POST(defines.EndpointCreateUser, usersCtrl.Create)
	r.GET(defines.EndpointGetUserByUUID, mwAuth.Check, usersCtrl.GetByUUID)
	r.POST(defines.EndpointCreateSession, sessionsCtrl.Create)
}
