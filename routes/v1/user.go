package v1

import (
	"tripat3k2/url_shortner/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
}
