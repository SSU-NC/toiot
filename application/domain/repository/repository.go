package repository

import "github.com/KumKeeHyun/PDK/application/domain/model"

type NodeRepository interface {
	GetAll() ([]model.Node, error)
	GetByUUID(string) (*model.Node, error)
	Create(*model.Node) error
	CreateNS(*model.NodeSensor) error
}

type SensorRepository interface {
	GetAll() ([]model.Sensor, error)
	GetByNodeUUID(string) ([]model.Sensor, error)
	GetByUUID(string) (*model.Sensor, error)
	GetValuesByUUID(string) ([]model.SensorValue, error)
	Create(*model.Sensor) error
	CreateValue(*model.SensorValue) error
}
