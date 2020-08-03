package adapter

import "github.com/KumKeeHyun/PDK/logic-core/domain/model"

type Node struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Location string `json:"location"`
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
		Name:  an.Name,
		Group: an.Location,
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
