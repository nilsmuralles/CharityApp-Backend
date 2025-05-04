package handlers

import (
	"log"
	"net/http"
	"time"

	"charityapp/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCampaignDonationsHandler(c *gin.Context, db *gorm.DB) {

	var results []*models.CampaignDonationsResponse

	err := db.Table("campania c").
		Select(`c.nombre AS campaign, c.fecha_fin as end_date,
			COALESCE(SUM(dm.monto), 0) AS monetary_total,
			COUNT(DISTINCT dnm.id) AS no_monetary_donations,
			COUNT(DISTINCT ad.id) AS total_donated_articles`).
		Joins("LEFT JOIN donacion_monetaria dm ON c.id = dm.id_campania").
		Joins("LEFT JOIN donacion_no_monetaria dnm ON c.id = dnm.id_campania").
		Joins("LEFT JOIN articulo_donado ad ON dnm.id = ad.id_donacion_no_monetaria").
		Group("c.id").
		Order("monetary_total DESC").
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

func GetTrendingDonationsHandler(c *gin.Context, db *gorm.DB) {
	startDate := c.DefaultQuery("start_date", "1970-01-01")
	endDate := c.DefaultQuery("end_date", time.Now().Format("2006-01-02"))

	var results []*models.TrendingDonationsResponse

	err := db.Table("donacion_monetaria dm").
		Select(`TO_CHAR(dm.fecha, 'YYYY-MM') AS month,
			SUM(dm.monto) AS total_donations,
			COUNT(DISTINCT dm.id_donador) AS unique_donors,
			dm.metodo_pago as pay_method`).
		Where("dm.fecha BETWEEN ? AND ?", startDate, endDate).
		Group("TO_CHAR(dm.fecha, 'YYYY-MM'), pay_method").
		Order("month").
		Scan(&results).Error

	if err != nil {
		log.Println("Error fetching donation trends:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}

func GetVolunteerParticipationHandler(c *gin.Context, db *gorm.DB) {
	var results []*models.VolunteerParticipationResponse

	err := db.Table("voluntario_voluntariado vv").
		Select(`c.nombre AS campaign,
			v.nombre AS volunteer,
			COUNT(vv.id_voluntariado) AS participations,
			SUM(EXTRACT(EPOCH FROM (vv.hora_fin - vv.hora_inicio)))/3600 AS total_hours`).
		Joins("JOIN voluntario v ON vv.id_voluntario = v.id").
		Joins("JOIN voluntariado vl ON vv.id_voluntariado = vl.id").
		Joins("JOIN campania c ON vl.id_campania = c.id").
		Group("campaign, volunteer").
		Order("total_hours DESC").
		Scan(&results).Error

	if err != nil {
		log.Println("Error fetching volunteer participation:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}

func GetDonorsDistributionHandler(c *gin.Context, db *gorm.DB) {
	var results []*models.DonorsDistributionResponse

	err := db.Table("donador d").
		Select(`d.categoria AS category,
			COUNT(DISTINCT d.id) AS total_donors,
			COUNT(DISTINCT dm.id) FILTER (WHERE dm.id IS NOT NULL) AS monetary_donors,
			COUNT(DISTINCT dnm.id) FILTER (WHERE dnm.id IS NOT NULL) AS no_monetary_donors,
			COALESCE(SUM(dm.monto), 0) AS total_amount`).
		Joins("LEFT JOIN donacion_monetaria dm ON d.id = dm.id_donador").
		Joins("LEFT JOIN donacion_no_monetaria dnm ON d.id = dnm.id_donador").
		Group("category").
		Scan(&results).Error

	if err != nil {
		log.Println("Error fetching donor distribution:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}

func GetCampaignEfficiencyHandler(c *gin.Context, db *gorm.DB) {
	var results []*models.CampaignEfficiencyResponse

	err := db.Table("campania c").
		Select(`c.nombre AS campaign,
			c.fecha_inicio AS init_date,
			c.fecha_fin AS end_date,
			COALESCE(SUM(dm.monto), 0) AS total_donations,
			COUNT(DISTINCT vv.id_voluntario) AS unique_volunteers,
			COUNT(DISTINCT vv.id) AS participations,
			COUNT(DISTINCT r.id) AS recognitions_awarded`).
		Joins("LEFT JOIN donacion_monetaria dm ON c.id = dm.id_campania").
		Joins("LEFT JOIN voluntariado vl ON c.id = vl.id_campania").
		Joins("LEFT JOIN voluntario_voluntariado vv ON vl.id = vv.id_voluntariado").
		Joins("LEFT JOIN reconocimiento r ON vv.id_voluntario = r.id_voluntario").
		Group("c.id").
		Order("init_date DESC").
		Scan(&results).Error

	if err != nil {
		log.Println("Error fetching campaign efficiency:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    results,
	})
}
