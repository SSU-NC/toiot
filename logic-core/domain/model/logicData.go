package model

import "time"

type LogicData struct {
	SID       string
	SName     string
	Values    map[string]float64
	NodeInfo  Node
	Timestamp time.Time
}

type LogicRing struct {
	Logic string                 `json:"logic"`
	Arg   map[string]interface{} `json:"arg"`
}

type ChainRequest struct {
	SID   string      `json:"s_id"`
	Name  string      `json:"name"`
	Rings []LogicRing `json:"rings"`
}
