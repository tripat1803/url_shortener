package main

import (
	"fmt"
	"tripat3k2/url_shortner/config"
	"tripat3k2/url_shortner/routes"
	"tripat3k2/url_shortner/utils"
)

func main() {
	config.GetConfigEnv()
	config.ConnectToDB()
	config.ConfigRouter()

	utils.MigrateModels(config.DB)

	routes.Init()

	port := fmt.Sprintf(":%s", config.Env.PORT)
	config.Router.Run(port)
}
