package adapter

import "github.com/KumKeeHyun/PDK/application/domain/model"

type Node struct {
	UUID     string         `json:"uuid"`
	Name     string         `json:"name"`
	Location string         `json:"location"`
	Sensors  []model.Sensor `json:"sensors"`
}

func ToNodes(n []model.Node) []Node {
	res := make([]Node, len(n))
	for i, node := range n {
		res[i] = ToNode(&node)
	}
	return res
}

func ToNode(n *model.Node) Node {
	return Node{
		UUID:     n.UUID,
		Name:     n.Name,
		Location: n.Location,
		Sensors:  make([]model.Sensor, 0),
	}
}
