package dblayer

import (
	"pdk/src/models"
)

type DBLayer interface {
	GetAllNodes() ([]models.Node, error)
	GetSensorsByNID(string) ([]models.Sensor, error)
	AddNode(models.Node) (models.Node, error)
	AddSensor(models.Sensor) (models.Sensor, error)
	AddNodeSensor(models.NodeSensor) (models.NodeSensor, error)
	AddSensorValue(models.SensorValue) (models.SensorValue, error)
	GetSensorValues(string) ([]string, error)
	GetAllSensors() ([]models.Sensor, error)
}
