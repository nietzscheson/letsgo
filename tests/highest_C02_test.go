package tests

import (
	"encoding/json"
	"letsgo/internal/app"
	"letsgo/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		panic(err)
	}
	return t
}

func TestHighestC02Route(t *testing.T) {
	router := app.SetupRouter()

	db := models.ConnectDatabase()

	defer db.Migrator().DropTable(&models.Sensor{})

	sensorData := []models.Sensor{
		{
			Time:     parseTime("2015-08-01T00:00:28Z"),
			Power:    0.0,
			Temp:     32,
			Humidity: 40,
			Light:    0,
			CO2:      973,
			Dust:     27.8,
		},
		{
			Time:     parseTime("2015-08-03T17:19:37Z"),
			Power:    0.572,
			Temp:     34,
			Humidity: 37,
			Light:    12,
			CO2:      900,
			Dust:     13.74,
		},
		{
			Time:     parseTime("2015-08-06T10:25:47Z"),
			Power:    -1,
			Temp:     33,
			Humidity: 36,
			Light:    16,
			CO2:      1362,
			Dust:     11.21,
		},
		{
			Time:     parseTime("2015-08-08T07:51:31Z"),
			Power:    0,
			Temp:     33,
			Humidity: 39,
			Light:    1,
			CO2:      1740,
			Dust:     14.6,
		},
	}

	for _, sensor := range sensorData {
		if result := db.Create(&sensor); result.Error != nil {
			t.Errorf("Failed to insert sensor data: %v", result.Error)
		}
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sensor/highest-c02", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "Expected HTTP status code to be 200")

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err, "Error decoding JSON")

	assert.Equal(t, "1740", response["highest_c02"], "Expected message to match '450'")
}
