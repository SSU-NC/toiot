package model

type Node struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type AppNode struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func ToNode(n AppNode) Node {
	return Node{
		Name:     n.Name,
		Location: n.Location,
	}
}

type Sensor struct {
	Name      string   `json:"name"`
	ValueList []string `json:"value_list"`
}

type AppSensor struct {
	UUID      string           `json:"uuid"`
	Name      string           `json:"name"`
	ValueList []AppSensorValue `json:"value_list"`
}

type AppSensorValue struct {
	SensorUUID string `json:"sensor_uuid"`
	ValueName  string `json:"value_name"`
	Index      int    `json:"index"`
}

func ToSensor(s AppSensor) Sensor {
	res := Sensor{
		Name:      s.Name,
		ValueList: make([]string, 0),
	}
	for _, v := range s.ValueList {
		res.ValueList = append(res.ValueList, v.ValueName)
	}
	return res
}
