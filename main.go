package main

import (
	"log"

	"campaign-app.local/config"
	"campaign-app.local/database"
	"campaign-app.local/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	database.ConnectDB()

	router := gin.New()
	router.Use(gin.Logger())

	routes.CampaignRoutes(router)

	log.Printf("Starting server on port %s", config.AppConfig.Port)
	router.Run(":" + config.AppConfig.Port)
}
