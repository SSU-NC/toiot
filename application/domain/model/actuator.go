package model

/*
액추에이터를 메모리에서 관리,
엑추에이터에게 보낼 [val, time] 리스트 전송하면 이거 따라서 진행
애플리케이션에서 바로 전송

싱크 추가할 때마다 싱크 주소 필요
로직조건 만족 하면 모든 싱크들에게 액추에이터 실행 메세지 전송
TCP받으면 

*/

type Actuator struct {
	ID             int             `json:"id" gorm:"primaryKey"`
	Name           string          `json:"name" gorm:"type:varchar(32);unique;not null"`
	ActuatorValues []ActuatorValue `json:"actuator_values" gorm:"foreignKey:ActuatorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Nodes          []Node          `json:"nodes" gorm:"many2many:has_actuators;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logics         []Logic         `json:"logics" gorm:"foreignKey:ActuatorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Actuator) TableName() string {
	return "actuator"
}

type ActuatorValue struct {
	ActuatorID int    `json:"sensor_id" gorm:"primaryKey"`
	ValueName  string `json:"value_name" gorm:"primaryKey;type:varchar(32)"`
	Index      int    `json:"index" gorm:"not null"`
}

func (ActuatorValue) TableName() string {
	return "actuator_values"
}
