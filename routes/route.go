package routes

import (
	"tripat3k2/url_shortner/config"
	"tripat3k2/url_shortner/controllers"
	v1 "tripat3k2/url_shortner/routes/v1"
)

func Init() {
	v1.V1()
	config.Router.GET("/:cipher", controllers.RedirectUrl)
}
