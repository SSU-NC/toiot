package model

type Topic struct {
	ID            int            `json:"id" gorm:"primary_key"`
	Name          string         `json:"name" gorm:"type:varchar(32);unique;not null"`
	Partitions    int            `json:"partitions"`
	Replications  int            `json:"replications"`
	Sinks         []Sink         `json:"sinks" gorm:"foreignkey:TopicID"`
	LogicServices []LogicService `json:"logic_services" gorm:"foreignkey:TopicID"`
}

func (Topic) TableName() string {
	return "topics"
}

type LogicService struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Addr    string `json:"addr" gorm:"type:varchar(32);not null"`
	TopicID int    `json:"topic_id" gorm:"not null"`
	Topic   Topic  `json:"topic" gorm:"foreignkey:TopicID"`
}

func (LogicService) TableName() string {
	return "logic_services"
}
