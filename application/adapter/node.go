package adapter

import "github.com/KumKeeHyun/toiot/application/domain/model"

type Node struct {
	ID       int            `json:"id"`
	Name     string         `json:"name"`
	Location Location       `json:"location"`
	SinkID   int            `json:"sink_id"`
	Sink     model.Sink     `json:"sink"`
	Sensors  []model.Sensor `json:"sensors"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Square struct {
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
	Up    float64 `json:"up"`
	Down  float64 `json:"down"`
}

type Page struct {
	Page int `json:"page"`
	Sink int `json:"sink_id"`
	Size int
}

func (p Page) GetOffset() int {
	return (p.Page - 1) * p.Size
}
