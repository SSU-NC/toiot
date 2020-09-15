package model

type Topic struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	Name          string         `json:"name" gorm:"type:varchar(32);unique;not null"`
	Partitions    int            `json:"partitions"`
	Replications  int            `json:"replications"`
	Sinks         []Sink         `json:"sinks" gorm:"foreignKey:TopicID"`
	LogicServices []LogicService `json:"logic_services" gorm:"foreignKey:TopicID"`
}

func (Topic) TableName() string {
	return "topics"
}

type LogicService struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Addr    string `json:"addr" gorm:"type:varchar(32);unique;not null"`
	TopicID int    `json:"topic_id" gorm:"not null"`
	Topic   Topic  `json:"topic" gorm:"foreignKey:TopicID"`
}

func (LogicService) TableName() string {
	return "logic_services"
}
