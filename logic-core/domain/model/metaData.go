package model

type Loc struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Node struct {
	Name     string `json:"name"`
	Group    string `json:"group"`
	Location Loc    `json:"location"`
}

type Sensor struct {
	Name       string   `json:"name"`
	ValueNames []string `json:"value_names"`
}
