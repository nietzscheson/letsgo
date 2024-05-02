package handler

import (
	"letsgo/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HighestC02Handler(c *gin.Context) {

	db := models.ConnectDatabase()

	var sensors []models.Sensor
	if err := db.Order("co2 desc").Find(&sensors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(sensors) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No sensors found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"highest_c02": strconv.Itoa(sensors[0].CO2),
	})
}

func HottestTemperature(c *gin.Context) {
	hottestTemperature := 38
	c.JSON(http.StatusOK, gin.H{
		"hottest_temperature": hottestTemperature,
	})
}

func HighestHumidity(c *gin.Context) {
	highestHumidity := 95
	c.JSON(http.StatusOK, gin.H{
		"highest_humidity": highestHumidity,
	})
}
