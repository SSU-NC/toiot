package adapter

import "github.com/KumKeeHyun/toiot/application/domain/model"

type Node struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Location Location       `json:"location"`
	SinkID   int            `json:"sink_id"`
	Sensors  []model.Sensor `json:"sensors"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
