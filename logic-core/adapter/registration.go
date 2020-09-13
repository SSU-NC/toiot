package adapter

import (
	"encoding/json"

	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
)

// INIT
// adapter : []Sink + Sink.Nodes + Sink.Nodes.Sensors + Sink.Nodes.Sesnors.Logics + Sink.Nodes.Sensors.SensorValues
// action : create Nodes, Sensors, Logics

// DeleteSink
// adapter : []Node + Node.Sink (Node.Sensors X)
// action : delete Nodes

// CreateNode
// adapter : Node + Node.Sink + Node.Sensors.SensorValues + Node.Sensors.Logics
// action : create Node, Sensors, Logics

// adapter.Node -> model.Node + []adapter.Sensor
// adapter.Sensor -> model.Sensor + []adpater.Logic
// []adapter.Logic -> []model.Logic

// DeleteNode
// adapter : Node (Node.Sink, Node.Sensors X)
// action : delete Node

// DeleteSensor
// adapter : Sensor + Sensor.Logics
// action : delete Sensor, Logics

// CreateLogic
// adapter : Logic (Logic.Sensor X)
// action : create Logic

// DeleteLogic
// adapter : Logic (Logic.Sensor X)
// action : delete Logic

type Logic struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Elems    string `json:"elems"`
	SensorID int    `json:"sensor_id"`
}

func LogicToModel(l *Logic) (model.Logic, error) {
	var elems []model.Element
	if err := json.Unmarshal([]byte(l.Elems), &elems); err != nil {
		return model.Logic{}, err
	} else {
		return model.Logic{
			ID:        l.ID,
			LogicName: l.Name,
			Elems:     elems,
			SensorID:  l.SensorID,
		}, nil
	}
}

func LogicsToModels(ll []Logic) []model.Logic {
	res := make([]model.Logic, 0, len(ll))
	for _, l := range ll {
		if ml, err := LogicToModel(&l); err == nil {
			res = append(res, ml)
		}
	}
	return res
}

type Sensor struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	SensorValues []SensorValue `json:"sensor_values"`
	Logics       []Logic       `json:"logics"`
}

type SensorValue struct {
	SensorID  int    `json:"sensor_id"`
	ValueName string `json:"value_name"`
	Index     int    `json:"index"`
}

func SensorToModel(s *Sensor) (model.Sensor, []Logic) {
	sv := make([]string, len(s.SensorValues))
	for i, v := range s.SensorValues {
		sv[i] = v.ValueName
	}
	return model.Sensor{
		Name:         s.Name,
		SensorValues: sv,
	}, s.Logics
}

type Node struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	LocLat  float64  `json:"lat"`
	LocLon  float64  `json:"lon"`
	SinkID  int      `json:"sink_id"`
	Sink    Sink     `json:"sink"`
	Sensors []Sensor `json:"sensors"`
}

func NodeToModel(n *Node, sn string) (model.Node, []Sensor) {
	return model.Node{
		Name: n.Name,
		Location: model.Location{
			Lat: n.LocLat,
			Lon: n.LocLon,
		},
		SinkName: sn,
	}, n.Sensors
}

type Sink struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Addr  string `json:"addr"`
	Nodes []Node `json:"nodes"`
}

type Topic struct {
	Name string `json:"name"`
}

type LogicService struct {
	Addr  string `json:"addr"`
	Topic Topic  `json:"topic"`
}
