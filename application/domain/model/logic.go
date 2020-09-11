package model

type Logic struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"type:varchar(32);unique;not null"`
	Elems    string `json:"elems" gorm:"type:varchar(128);not null"`
	SensorID int    `json:"sensor_id" gorm:"not null"`
	Sensor   Sensor `json:"sensor" gorm:"foreignkey:SensorID"`
}

func (Logic) TableName() string {
	return "logics"
}
