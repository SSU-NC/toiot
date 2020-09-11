package usecase

import (
	"github.com/KumKeeHyun/toiot/application/domain/model"
)

// for ui registration
type RegistUsecase interface {
	GetSinks() ([]model.Sink, error)
	RegistSink(s *model.Sink) error
	UnregistSink(s *model.Sink) error

	GetNodes() ([]model.Node, error)
	RegistNode(n *model.Node) error
	UnregistNode(n *model.Node) error

	GetSensors() ([]model.Sensor, error)
	RegistSensor(s *model.Sensor) error
	UnregistSensor(s *model.Sensor) error

	GetLogics() ([]model.Logic, error)
	RegistLogic(l *model.Logic) error
	UnregistLogic(l *model.Logic) error

	GetLogicServices() ([]model.LogicService, error)
	UnregistLogicService(l *model.LogicService) error

	GetTopics() ([]model.Topic, error)
	RegistTopic(t *model.Topic) error
	UnregistTopic(t *model.Topic) error
}

// for event channel
type EventUsecase interface {
	RegistLogicService(l *model.LogicService) error
}
