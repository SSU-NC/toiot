package usecase

import (
	"github.com/KumKeeHyun/PDK/application/domain/model"
)

type SinkUsecase interface {
	GetAllSinks() ([]model.Sink, error)
	GetAllSinksWithNodes() ([]model.Sink, error)
	GetSinkByID(uint) (*model.Sink, error)
	GetSinkByIDWithNodes(uint) (*model.Sink, error)
	RegisterSink(*model.Sink) (*model.Sink, error)
	DeleteSink(*model.Sink) error
}

type NodeUsecase interface {
	GetAllNodes() ([]model.Node, error)
	GetAllNodesWithSensors() ([]model.Node, error)
	GetAllNodesWithSensorsWithValues() ([]model.Node, error)
	GetNodeByUUID(string) (*model.Node, error)
	GetNodeByUUIDWithSensors(string) (*model.Node, error)
	GetNodesBySinkID(uint) ([]model.Node, error)
	RegisterNode(*model.Node) (*model.Node, error)
	DeleteNode(*model.Node) (*model.Node, error)
}

type SensorUsecase interface {
	GetAllSensors() ([]model.Sensor, error)
	GetAllSensorsWithValues() ([]model.Sensor, error)
	RegisterSensor(*model.Sensor) (*model.Sensor, error)
	DeleteSensor(*model.Sensor) (*model.Sensor, error)
}
