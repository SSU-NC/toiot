package usecase

import (
	"github.com/KumKeeHyun/PDK/application/domain/model"
	"github.com/KumKeeHyun/PDK/application/interface/presenter"
)

type NodeUsecase interface {
	GetAllNodes() ([]presenter.Node, error)
	GetRegister() ([]model.Node, error)
	RegisterNode(*presenter.Node) (*model.Node, error)
	DeleteNode(*presenter.Node) (*model.Node, error)
}

type SensorUsecase interface {
	GetAllSensors() ([]model.Sensor, error)
	GetRegister() ([]model.Sensor, error)
	RegisterSensor(*model.Sensor) (*model.Sensor, error)
	DeleteSensor(*model.Sensor) (*model.Sensor, error)
}
