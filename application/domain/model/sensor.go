package model

type Sensor struct {
	ID           int           `json:"id" gorm:"primaryKey"`
	Name         string        `json:"name" gorm:"type:varchar(32);unique;not null"`
	SensorValues []SensorValue `json:"sensor_values" gorm:"foreignKey:SensorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Nodes        []Node        `json:"nodes" gorm:"many2many:has_sensors;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logics       []Logic       `json:"logics" gorm:"foreignKey:SensorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Sensor) TableName() string {
	return "sensors"
}

type SensorValue struct {
	SensorID  int    `json:"sensor_id" gorm:"primaryKey"`
	ValueName string `json:"value_name" gorm:"primaryKey;type:varchar(32)"`
	Index     int    `json:"index" gorm:"not null"`
}

func (SensorValue) TableName() string {
	return "sensor_values"
}
