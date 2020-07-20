package model

import "github.com/segmentio/ksuid"

type Node struct {
	UUID     string `gorm:"primary_key;type:char(27);not null;"`
	Name     string `gorm:"type:varchar(32);not null"`
	Location string `gorm:"type:varchar(64)"`
	//Sensors  []*Sensor `gorm:"many2many:node_sensors"`
}

func NewNode(name, loc string) Node {
	return Node{
		UUID:     ksuid.New().String(),
		Name:     name,
		Location: loc,
	}
}

func (Node) TableName() string {
	return "nodes"
}

type NodeSensor struct {
	NodeUUID   string `gorm:"primary_key;type:char(27)"`
	SensorUUID string `gorm:"primary_key;type:char(27)"`
}

func (NodeSensor) TableName() string {
	return "node_sensors"
}
