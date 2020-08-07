package usecase

import (
	"github.com/KumKeeHyun/PDK/application/adapter"
	"github.com/KumKeeHyun/PDK/application/domain/model"
)

type NodeUsecase interface {
	GetAllNodes() ([]adapter.Node, error)
	GetRegister() ([]model.Node, error)
	RegisterNode(*adapter.Node) (*model.Node, error)
	DeleteNode(*adapter.Node) (*model.Node, error)
}

type SensorUsecase interface {
	GetAllSensors() ([]model.Sensor, error)
	GetRegister() ([]model.Sensor, error)
	RegisterSensor(*model.Sensor) (*model.Sensor, error)
	DeleteSensor(*model.Sensor) (*model.Sensor, error)
}
