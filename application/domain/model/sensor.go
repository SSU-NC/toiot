package model

type Sensor struct {
	ID           int           `json:"id" gorm:"primary_key"`
	Name         string        `json:"name" gorm:"type:varchar(32);unique;not null"`
	SensorValues []SensorValue `json:"sensor_values" gorm:"foreignkey:SensorID"`
	Nodes        []Node        `json:"nodes" gorm:"many2many:has_sensors"`
}

func (Sensor) TableName() string {
	return "sensors"
}

type SensorValue struct {
	SensorID  int    `json:"sensor_id" gorm:"primary_key"`
	ValueName string `json:"value_name" gorm:"primary_key;type:varchar(32)"`
	Index     int    `json:"index" gorm:"not null"`
}

func (SensorValue) TableName() string {
	return "sensor_values"
}
