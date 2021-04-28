package model

type Actuator struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(32);unique;not null"`
}

func (Actuator) TableName() string {
	return "actuators"
}
