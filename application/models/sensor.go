package models

import (
	"github.com/segmentio/ksuid"
)

type Sensor struct {
	UUID        string   `json:"uuid" gorm:"primary_key;type:char(27);not null;"`
	Name        string   `json:"name" gorm:"type:varchar(32);not null"`
	NumOfValues int      `json:"num_of_values" gorm:"not null"`
	ValueNames  []string `json:"value_names" gorm:"-"`
}

func (Sensor) TableName() string {
	return "sensors"
}

func NewSensor() Sensor {
	s := Sensor{}
	s.UUID = ksuid.New().String()
	return s
}

type SensorValue struct {
	SensorUUID string `json:"sensor_uuid" gorm:"primary_key;type:char(27);not null"`
	ValueName  string `json:"value_name" gorm:"primary_key;type:varchar(32);not null"`
	Index      int    `json:"index" gorm:"not null"`
}

func (SensorValue) TableName() string {
	return "sensor_values"
}
