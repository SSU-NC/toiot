package adapter

import "github.com/seheee/PDK/application/domain/model"

type Node struct {
	UUID     string         `json:"uuid"`
	Name     string         `json:"name"`
	Group    string         `json:"group"`
	Location Loc            `json:"location"`
	SinkID   uint           `json:"sink_id"`
	Sensors  []model.Sensor `json:"sensors"`
}

type Loc struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func ModelsToNodes(n []model.Node) []Node {
	res := make([]Node, len(n))
	for i, node := range n {
		res[i] = ModelToNode(&node)
	}
	return res
}

func ModelToNode(n *model.Node) Node {
	return Node{
		UUID:  n.UUID,
		Name:  n.Name,
		Group: n.Group,
		Location: Loc{
			Lat: n.LocLat,
			Lon: n.LocLon,
		},
		SinkID:  n.SinkID,
		Sensors: n.Sensors,
	}
}

func NodesToModels(n []Node) []model.Node {
	res := make([]model.Node, len(n))
	for i, node := range n {
		res[i] = NodeToModel(&node)
	}
	return res
}

func NodeToModel(n *Node) model.Node {
	return model.Node{
		UUID:    n.UUID,
		Name:    n.Name,
		Group:   n.Group,
		LocLat:  n.Location.Lat,
		LocLon:  n.Location.Lon,
		SinkID:  n.SinkID,
		Sensors: n.Sensors,
	}
}
