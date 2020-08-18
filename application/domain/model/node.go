package model

import "github.com/rs/xid"

type Node struct {
	UUID    string   `json:"uuid" gorm:"primary_key;type:char(20);not null;"`
	Name    string   `json:"name" gorm:"type:varchar(32);unique;not null"`
	Group   string   `json:"location" gorm:"type:varchar(64)"`
	LocLat  float64  `json:"lat"`
	LocLon  float64  `json:"lon"`
	SinkID  uint     `json:"sink_id" gorm:"not null"`
	Sensors []Sensor `json:"sensors" gorm:"foreignkey:UUID"`
}

func NewNode(name, grp string, lat, lon float64, sinkID uint) Node {
	return Node{
		UUID:   xid.New().String(),
		Name:   name,
		Group:  grp,
		LocLat: lat,
		LocLon: lon,
		SinkID: sinkID,
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
