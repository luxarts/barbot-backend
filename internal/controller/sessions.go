package controller

import (
	"barbot/internal/domain"
	"barbot/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionsController interface {
	Create(ctx *gin.Context)
}

type sessionsController struct {
	svc service.SessionsService
}

func NewSessionsController(svc service.SessionsService) SessionsController {
	return &sessionsController{svc: svc}
}

func (c *sessionsController) Create(ctx *gin.Context) {
	var body domain.SessionDTO
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("wrong body"))
		return
	}

	token := c.svc.Create(body)

	if token == nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("wrong credentials"))
		return
	}

	ctx.JSON(http.StatusOK, domain.SessionResponse{Token: *token})
}
