package middlewares

import (
	"net/http"
	"strings"
	"tripat3k2/url_shortner/utils"

	"github.com/gin-gonic/gin"
)

type AuthHeader struct {
	Authorization string
}

func VerifyAuth(ctx *gin.Context) {
	header := &AuthHeader{}
	err := ctx.BindHeader(header)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Token not found"})
		return
	}

	if !strings.HasPrefix(header.Authorization, "Bearer ") {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid token"})
		return
	}

	token := header.Authorization[7:]
	claims, tokenErr := utils.VerifyToken(token)
	if tokenErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error validating token", "error": tokenErr})
		return
	}

	ctx.Set("userId", claims.UserId)
	ctx.Next()
}
