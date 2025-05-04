package handlers

import (
	"log"
	"net/http"

	"charityapp/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCampaignDonationsHandler(c *gin.Context, db *gorm.DB) {

	var results []*models.CampaignDonationsResponse

	err := db.Table("campania c").
		Select(`c.nombre AS campaña,
			COALESCE(SUM(dm.monto), 0) AS total_monetario,
			COUNT(DISTINCT dnm.id) AS cantidad_donaciones_no_monetarias,
			COUNT(DISTINCT ad.id) AS total_articulos_donados`).
		Joins("LEFT JOIN donacion_monetaria dm ON c.id = dm.id_campania").
		Joins("LEFT JOIN donacion_no_monetaria dnm ON c.id = dnm.id_campania").
		Joins("LEFT JOIN articulo_donado ad ON dnm.id = ad.id_donacion_no_monetaria").
		Group("c.id").
		Order("total_monetario DESC").
		Scan(&results).Error

	if err != nil {
		log.Println("Error fetching donations by campaign:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}
