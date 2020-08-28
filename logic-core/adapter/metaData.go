package adapter

import "github.com/KumKeeHyun/PDK/logic-core/domain/model"

type MetaInfo struct {
	NInfo []Node   `json:"node_info"`
	SInfo []Sensor `json:"sensor_info"`
}

type Node struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Group    string `json:"group"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}

type Sensor struct {
	UUID      string        `json:"uuid"`
	Name      string        `json:"name"`
	ValueList []SensorValue `json:"value_list"`
}

type SensorValue struct {
	SensorUUID string `json:"sensor_uuid"`
	ValueName  string `json:"value_name"`
	Index      int    `json:"index"`
}

func AppToNode(an *Node) model.Node {
	return model.Node{
		Name:     an.Name,
		Group:    an.Group,
		Location: an.Location,
	}
}

func AppToSensor(as *Sensor) model.Sensor {
	s := model.Sensor{
		Name:       as.Name,
		ValueNames: make([]string, len(as.ValueList)),
	}
	for i, v := range as.ValueList {
		s.ValueNames[i] = v.ValueName
	}
	return s
}
