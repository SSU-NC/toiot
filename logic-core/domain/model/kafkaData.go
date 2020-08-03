package model

import "time"

type SensorData struct {
	NID       string    `json:"nid"`
	Values    []float64 `json:"values"`
	Timestamp time.Time `json:"timestamp"`
}

type KafkaData struct {
	Key   string     `json:"key"`
	Value SensorData `json:"value"`
}
