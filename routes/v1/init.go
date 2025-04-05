package v1

import "tripat3k2/url_shortner/config"

func V1() {
	routerGroup := config.Router.Group("/v1")

	{
		user := routerGroup.Group("/user")
		UserRoutes(user)
	}
}
