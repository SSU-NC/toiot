package model

import "time"

type LogicData struct {
	SID       string             `json:"s_id"`
	SName     string             `json:"name"`
	Values    map[string]float64 `json:"values"`
	NodeInfo  Node               `json:"node"`
	Timestamp time.Time          `json:"timestamp"`
}

type LogicRing struct {
	Logic string                 `json:"logic"`
	Arg   map[string]interface{} `json:"arg"`
}

type ChainRequest struct {
	SID   string      `json:"s_id"`
	Name  string      `json:"name"`
	Rings []LogicRing `json:"rings"`
}

type RingRequest struct {
	Sensor string `json:"sensor_uuid"`
	LogicName string `json:"logic_name"`
	Logic []struct {
		Elem string `json:"elem"`
		Arg map[string]interface{} `json:"arg"`
	} `json:"logic"`
}