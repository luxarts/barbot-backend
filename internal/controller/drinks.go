package controller

import (
	"barbot/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DrinksController interface {
	GetAll(ctx *gin.Context)
}
type drinksController struct {
	srv service.DrinksService
}

func NewDrinksController(srv service.DrinksService) DrinksController {
	return &drinksController{srv: srv}
}

func (c *drinksController) GetAll(ctx *gin.Context) {
	mixedDrinks := c.srv.GetAll()

	ctx.JSON(http.StatusOK, mixedDrinks)
}
