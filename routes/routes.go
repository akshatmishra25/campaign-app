package routes

import (
	"campaign-app.local/controllers"

	"github.com/gin-gonic/gin"
)

func CampaignRoutes(router *gin.Engine) {
	router.GET("/campaigns", controllers.GetCampaigns)
	router.GET("/campaigns/:id", controllers.GetCampaignByID)
	router.POST("/campaigns", controllers.CreateCampaign)
	router.PUT("/campaigns/:id", controllers.UpdateCampaign)
	router.DELETE("/campaigns/:id", controllers.DeleteCampaign)
}