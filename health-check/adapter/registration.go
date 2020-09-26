package adapter

type Node struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Group    string `json:"group"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
	// SinkID uint `json:"sink_id"`
}

type Sink struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Addr string `json:"addr"`
}
