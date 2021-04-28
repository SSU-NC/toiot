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
	Left  float64 `form:"left" json:"left"`
	Right float64 `form:"right" json:"right"`
	Up    float64 `form:"up" json:"up"`
	Down  float64 `form:"down" json:"down"`
}

func (sq Square) IsBinded() bool {
	if sq.Left != 0 || sq.Right != 0 || sq.Up != 0 || sq.Down != 0 {
		return true
	}
	return false
}

type Page struct {
	Page int `form:"page" json:"page"`
	Sink int `form:"sink" json:"sink"`
	Size int `form:"size" json:"size"`
}

func (p Page) IsBinded() bool {
	if p.Page != 0 {
		return true
	}
	return false
}

func (p Page) GetOffset() int {
	return (p.Page - 1) * p.Size
}

type SinkPage struct {
	Sinks []model.Sink `json:"sinks"`
	Pages int          `json:"pages"`
}

type NodePage struct {
	Nodes []model.Node `json:"nodes"`
	Pages int          `json:"pages"`
}

type SensorPage struct {
	Sensors []model.Sensor `json:"sensors"`
	Pages   int            `json:"pages"`
}

type SinkAddr struct {
	Sid  int    `json:"sid"`
	Addr string `json:"addr"`
}
