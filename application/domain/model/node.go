package model

type Sink struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"type:varchar(32);unique;not null"`
	Addr    string `json:"addr" gorm:"type:varchar(32);not null"`
	TopicID int    `json:"topic_id" gorm:"not null"`
	Topic   Topic  `json:"topic" gorm:"foreignKey:TopicID"`
	Nodes   []Node `json:"nodes" gorm:"foreignKey:SinkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}


func (Sink) TableName() string {
	return "sinks"
}

type Node struct {
	ID      int      `json:"id" gorm:"primaryKey"`
	Name    string   `json:"name" gorm:"type:varchar(32);unique;not null"`
	LocLat  float64  `json:"lat"`
	LocLon  float64  `json:"lng"`
	SinkID  int      `json:"sink_id" gorm:"not null"`
	Sink    Sink     `json:"sink" gorm:"foreignKey:SinkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Sensors []Sensor `json:"sensors" gorm:"many2many:has_sensors;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Node) TableName() string {
	return "nodes"
}
