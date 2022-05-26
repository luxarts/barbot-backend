package controller

import (
	"barbot/internal/defines"
	"barbot/internal/domain"
	"barbot/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionsController interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
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

func (c *sessionsController) Get(ctx *gin.Context) {
	userID, exists := ctx.Get(defines.ParamUserID)
	if !exists {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("failed to obtain payload"))
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"userID": userID.(string)})
}
