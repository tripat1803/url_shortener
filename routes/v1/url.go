package v1

import (
	"tripat3k2/url_shortner/controllers"
	"tripat3k2/url_shortner/middlewares"

	"github.com/gin-gonic/gin"
)

func UrlRoutes(router *gin.RouterGroup) {
	router.POST("/create", middlewares.VerifyAuth, controllers.CreateShortUrl)
}
