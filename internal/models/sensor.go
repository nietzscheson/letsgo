package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sensor struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Time     time.Time
	Power    float64 `gorm:"default:0.0"`
	Temp     int
	Humidity int
	Light    int
	CO2      int
	Dust     float64 `gorm:"default:0.0"`
}

func (sensor *Sensor) BeforeCreate(tx *gorm.DB) (err error) {
	sensor.ID = uuid.New().String()
	return
}
