package model

type Node struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
	SinkName string   `json:"sink_name"`
	Nid      int      `json:"nid"`
	Sid      int      `json:"sid"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Sensor struct {
	Name         string   `json:"name"`
	SensorValues []string `json:"sensor_values"`
}

type Sink struct {
	// Sid		 int 	  `json:"sid"`
	Addr string `json:"addr"`
}
type Nodeinfo struct {
	SinkID int `json:"sink_id"`
}
