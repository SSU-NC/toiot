package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogicData struct {
	SID       string             `json:"s_id"`
	SName     string             `json:"name"`
	Values    map[string]float64 `json:"values"`
	NodeInfo  Node               `json:"node"`
	Timestamp time.Time          `json:"timestamp"`
}

type RingRequest struct {
	Sensor string `json:"sensor_uuid"`
	LogicName string `json:"logic_name"`
	Logic []struct {
		Elem string `json:"elem"`
		Arg map[string]interface{} `json:"arg"`
	} `json:"logic"`
}

type Ring struct {
	Id primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Sensor string `json:"sensor_uuid"`
	LogicName string `json:"logic_name"`
	Logic []struct {
		Elem string `json:"elem"`
		Arg map[string]interface{} `json:"arg"`
	} `json:"logic"`
}