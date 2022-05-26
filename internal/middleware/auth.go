package middleware

import (
	"barbot/internal/defines"
	"barbot/internal/utils/jwt"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type AuthMiddleware interface {
	Check(c *gin.Context)
}

type authMiddleware struct {
}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

func (mw *authMiddleware) Check(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New("missing Authorization token"))
		return
	}

	bearerTokenSplit := strings.Split(bearerToken, " ")

	if len(bearerTokenSplit) != 2 || bearerTokenSplit[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New("invalid Authorization token"))
		return
	}

	secret := os.Getenv(defines.EnvJWTSecret)
	if !jwt.Verify(bearerTokenSplit[1], secret) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New("invalid signature"))
		return
	}

	// Set userID in context
	payload, err := jwt.GetPayload(bearerTokenSplit[1])

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	userID := payload.Subject

	ctx.Set(defines.ParamUserID, userID)
}
