package controller

import (
	"barbot/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MixedDrinksController interface {
	GetAll(ctx *gin.Context)
}
type mixedDrinksController struct {
	srv service.MixedDrinksService
}

func NewMixedDrinksController(srv service.MixedDrinksService) MixedDrinksController {
	return &mixedDrinksController{srv: srv}
}

func (c *mixedDrinksController) GetAll(ctx *gin.Context) {
	mixedDrinks := c.srv.GetAll()

	ctx.JSON(http.StatusOK, mixedDrinks)
}
