package models

import (
	"github.com/segmentio/ksuid"
)

type Node struct {
	UUID     string   `json:"uuid" gorm:"primary_key;type:char(27);not null;"`
	Name     string   `json:"name" gorm:"type:varchar(32);not null"`
	Location string   `json:"location" gorm:"type:varchar(64)"`
	Sensors  []Sensor `json:"sensors" gorm:"-"`
}

func (Node) TableName() string {
	return "nodes"
}

func NewNode() Node {
	n := Node{}
	n.UUID = ksuid.New().String()
	return n
}

type NodeSensor struct {
	NodeUUID   string `json:"node_uuid" gorm:"primary_key;type:char(27)"`
	SensorUUID string `json:"sensor_uuid" gorm:"primary_key;type:char(27)"`
}

func (NodeSensor) TableName() string {
	return "node_sensors"
}
