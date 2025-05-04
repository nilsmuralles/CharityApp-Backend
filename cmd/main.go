package main

import (
	"charityapp/app/handlers"
	"charityapp/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	db := database.ConnectToDatabase()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "¡Backend conectado con éxito!",
		})
	})
	api := router.Group("/api")
	{
		api.GET("/campaign-donations", func(c *gin.Context) {
			handlers.GetCampaignDonationsHandler(c, db)
		})
		api.GET("/trending-donations", func(c *gin.Context) {
			handlers.GetTrendingDonationsHandler(c, db)
		})
		api.GET("/volunteer-participation", func(c *gin.Context) {
			handlers.GetVolunteerParticipationHandler(c, db)
		})
		api.GET("/donors-distribution", func(c *gin.Context) {
			handlers.GetDonorsDistributionHandler(c, db)
		})
		api.GET("/campaign-efficiency", func(c *gin.Context) {
			handlers.GetCampaignEfficiencyHandler(c, db)
		})
	}

	router.Run(":8080")
}
