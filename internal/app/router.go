package app

import (
	"letsgo/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("ping", handler.PingHandler)
	router.GET("sensor/highest-c02", handler.HighestC02Handler)
	router.GET("sensor/hottest-temperature", handler.HottestTemperature)
	router.GET("sensor/highest-humidity", handler.HighestHumidity)

	return router
}
