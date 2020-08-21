package model

import "github.com/rs/xid"

type Sensor struct {
	UUID      string        `json:"uuid" gorm:"primary_key;type:char(20);not null;"`
	Name      string        `json:"name" gorm:"type:varchar(32);unique;not null"`
	ValueList []SensorValue `json:"value_list" gorm:"foreignkey:SensorUUID"`
}

func NewSensor(name string) Sensor {
	return Sensor{
		UUID: xid.New().String(),
		Name: name,
	}
}

func (Sensor) TableName() string {
	return "sensors"
}

type SensorValue struct {
	SensorUUID string `json:"sensor_uuid" gorm:"primary_key;type:char(20);not null"`
	ValueName  string `json:"value_name" gorm:"primary_key;type:varchar(32);not null"`
	Index      int    `json:"index" gorm:"not null"`
}

func (SensorValue) TableName() string {
	return "sensor_values"
}
