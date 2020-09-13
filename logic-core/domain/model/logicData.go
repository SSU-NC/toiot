package model

import (
	"time"
)

type LogicData struct {
	SensorID   int                `json:"sensor_id"`
	SensorName string             `json:"sensor_name"`
	Values     map[string]float64 `json:"values"`
	Node       Node               `json:"node"`
	Timestamp  time.Time          `json:"timestamp"`
}

type Logic struct {
	ID        int       `json:"id"`
	LogicName string    `json:"logic_name"`
	Elems     []Element `json:"elems"`
	SensorID  int       `json:"sensor_id"`
}

type Element struct {
	Elem string                 `json:"elem"`
	Arg  map[string]interface{} `json:"arg"`
}
