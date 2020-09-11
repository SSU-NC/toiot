package adapter

type Logic struct {
	ID        int       `json:"id"`
	SensorID  int       `json:"sensor_id"`
	LogicName string    `json:"logic_name"`
	Elems     []Element `json:"elems"`
}

type Element struct {
	Elem string                 `json:"elem"`
	Arg  map[string]interface{} `json:"arg"`
}
