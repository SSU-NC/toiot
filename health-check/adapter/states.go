package adapter

type States struct {
	Timestamp string      `json:"timestamp"`
	State     []NodeState `json:"state"`
}

type NodeState struct {
	NodeID int  `json:"nid"`
	State  bool `json:"state"`
}

// type HealthInfo struct {
// 	UUID  string `json:"n_uuid"`
// 	State int    `json:"state"`
// }
