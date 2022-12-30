package models

import "time"

type SensorDataModel struct {
	ID          int
	Temperature float32
	Humidity    float32
	Pressure    float32
	VOC         float32
	UV          float32
	Light       float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PicoId      time.Time
}
