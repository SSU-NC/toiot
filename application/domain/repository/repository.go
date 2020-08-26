package repository

import "github.com/KumKeeHyun/PDK/application/domain/model"

type SinkRepository interface {
	GetAll() ([]model.Sink, error)
	GetAllWithNodes() ([]model.Sink, error)
	GetByID(uint) (*model.Sink, error)
	GetByIDWithNodes(uint) (*model.Sink, error)
	Create(*model.Sink) error
	Delete(*model.Sink) error
}

type NodeRepository interface {
	GetAll() ([]model.Node, error)
	GetByUUIDs([]string) ([]model.Node, error)
	GetByUUID(string) (*model.Node, error)
	GetBySinkID(uint) ([]model.Node, error)
	Create(*model.Node) error
	Delete(*model.Node) error
	CreateNS(*model.NodeSensor) error
}

type SensorRepository interface {
	// SELECT * FROM sensors
	GetAll() ([]model.Sensor, error)

	// SELECT * FROM sensors JOIN nodesensors(s_uuid) WHERE n_uuid = string
	GetByNodeUUID(string) ([]model.Sensor, error)

	// SELECT * FROM sensors JOIN sensor_values(sensor_uuid) JOIN nodesensors(s_uuid) WHERE n_uuid = string
	GetByNodeUUIDWithValues(string) ([]model.Sensor, error)

	// SELECT * FROM sensors WHERE uuid = string
	GetByUUID(string) (*model.Sensor, error)

	// SELECT * FROM sensors JOIN sensor_values(sensor_uuid) WHERE uuid = string
	GetByUUIDWithValues(string) (*model.Sensor, error)

	// SELECT * FROM sensor_values WHERE sensor_uuid = string
	GetValuesByUUID(string) ([]model.SensorValue, error)
	Create(*model.Sensor) error
	Delete(*model.Sensor) error
	CreateValue(*model.SensorValue) error
}
