package wsClient

const (
	Init = iota
	NewNode
	UpdateNode
	DeleteNode
	NewSensor
	DeleteSensor
)

type KafkaMessage struct {
	Type int                    `json:"type"`
	Msg  map[string]interface{} `json:"message"`
}

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

type RegisterInit struct {
	NodeInfo   []Node   `json:"node_info"`
	SensorInfo []Sensor `json:"sensor_info"`
}

type RegisterInfo struct {
	NodeInfo   map[string]Node
	SensorInfo map[string]Sensor
}
