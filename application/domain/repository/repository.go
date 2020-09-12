package repository

import (
	"github.com/KumKeeHyun/toiot/application/domain/model"
)

type SinkRepo interface {
	FindsWithTopic() ([]model.Sink, error)
	FindByIDWithNodesSensorsValuesTopic(id int) (*model.Sink, error)
	Create(*model.Sink) error
	Delete(*model.Sink) error
}

type NodeRepo interface {
	FindsWithSensorsValues() ([]model.Node, error)
	Create(*model.Node) error
	Delete(*model.Node) error
}

type SensorRepo interface {
	FindsWithValues() ([]model.Sensor, error)
	Create(*model.Sensor) error
	Delete(*model.Sensor) error
}

type LogicRepo interface {
	FindsWithSensorValues() ([]model.Logic, error)
	Create(*model.Logic) error
	Delete(*model.Logic) error
}

type LogicServiceRepo interface {
	Finds() ([]model.LogicService, error)
	FindsWithTopic() ([]model.LogicService, error)
	FindsByTopicID(TopicID int) ([]model.LogicService, error)
	Create(*model.LogicService) error
	Delete(*model.LogicService) error
}

type TopicRepo interface {
	FindsWithLogicService() ([]model.Topic, error)
	Create(*model.Topic) error
	Delete(*model.Topic) error
}
