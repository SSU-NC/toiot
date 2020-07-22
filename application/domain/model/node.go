package model

import "github.com/rs/xid"

type Node struct {
	UUID     string `json:"uuid" gorm:"primary_key;type:char(20);not null;"`
	Name     string `json:"name" gorm:"type:varchar(32);not null"`
	Location string `json:"location" gorm:"type:varchar(64)"`
	//Sensors  []*Sensor `gorm:"many2many:node_sensors"`
}

func NewNode(name, loc string) Node {
	return Node{
		UUID:     xid.New().String(),
		Name:     name,
		Location: loc,
	}
}

func (Node) TableName() string {
	return "nodes"
}

type NodeSensor struct {
	NodeUUID   string `gorm:"primary_key;type:char(20)"`
	SensorUUID string `gorm:"primary_key;type:char(20)"`
}

func (NodeSensor) TableName() string {
	return "node_sensors"
}
