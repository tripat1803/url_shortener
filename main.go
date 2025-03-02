package main

import (
	"fmt"
	"tripat3k2/url_shortner/config"
	"tripat3k2/url_shortner/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectToDB()
	utils.MigrateModels(db)

	router := gin.Default()

	env := config.GetConfigEnv()
	port := fmt.Sprintf(":%s", env.PORT)
	router.Run(port)
}
