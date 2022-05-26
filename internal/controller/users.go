package controller

import (
	"barbot/internal/defines"
	"barbot/internal/domain"
	"barbot/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController interface {
	Create(ctx *gin.Context)
	GetByUUID(ctx *gin.Context)
}

type usersController struct {
	svc service.UsersService
}

func NewUsersController(svc service.UsersService) UsersController {
	return &usersController{svc: svc}
}

func (c *usersController) Create(ctx *gin.Context) {
	var user domain.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("wrong body"))
		return
	}

	userDTO, err := c.svc.Create(user)
	if err != nil {
		if err == defines.ErrEmailAlreadyUsed {
			ctx.JSON(http.StatusConflict, map[string]string{"message": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, userDTO)
}
func (c *usersController) GetByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	ctxUserUUID, _ := ctx.Get(defines.ParamUserID)

	if ctxUserUUID != uuid {
		ctx.AbortWithError(http.StatusForbidden, errors.New("uuid not allowed"))
		return
	}

	userDTO := c.svc.GetByUUID(uuid)
	if userDTO == nil {
		ctx.AbortWithError(http.StatusNotFound, errors.New("user not found"))
		return
	}

	ctx.JSON(http.StatusOK, userDTO)
}
