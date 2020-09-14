package model

type Logic struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(32);unique;not null"`
	Elems    string `json:"elems" gorm:"type:text;not null"`
	SensorID int    `json:"sensor_id" gorm:"not null"`
	Sensor   Sensor `json:"sensor" gorm:"foreignKey:SensorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Logic) TableName() string {
	return "logics"
}
