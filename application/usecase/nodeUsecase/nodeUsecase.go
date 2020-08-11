package nodeUsecase

import (
	"github.com/seheee/PDK/application/domain/model"
	"github.com/seheee/PDK/application/domain/repository"
	"github.com/seheee/PDK/application/interface/presenter"
)

type nodeUsecase struct {
	nr repository.NodeRepository
	sr repository.SensorRepository
}

func NewNodeUsecase(nr repository.NodeRepository, sr repository.SensorRepository) *nodeUsecase {
	return &nodeUsecase{
		nr: nr,
		sr: sr,
	}
}

func (nu *nodeUsecase) GetAllNodes() ([]presenter.Node, error) {
	ns, err := nu.nr.GetAll()
	if err != nil {
		return nil, err
	}
	nodes := presenter.ToNodes(ns)
	for i := range nodes {
		nodes[i].Sensors, err = nu.sr.GetByNodeUUID(nodes[i].UUID)
		if err != nil {
			return nil, err
		}
		for j := range nodes[i].Sensors {
			nodes[i].Sensors[j].ValueList, err = nu.sr.GetValuesByUUID(nodes[i].Sensors[j].UUID)
			if err != nil {
				return nil, err
			}
		}
	}
	return nodes, nil
}

func (nu *nodeUsecase) GetRegister() ([]model.Node, error) {
	nodes, err := nu.nr.GetAll()
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (nu *nodeUsecase) RegisterNode(n *presenter.Node) (*model.Node, error) {
	newNode := model.NewNode(n.Name, n.Location)
	if err := nu.nr.Create(&newNode); err != nil {
		return nil, err
	}
	for _, s := range n.Sensors {
		ns := &model.NodeSensor{
			NodeUUID:   newNode.UUID,
			SensorUUID: s.UUID,
		}
		if err := nu.nr.CreateNS(ns); err != nil {
			return nil, err
		}
	}
	return &newNode, nil
}

func (nu *nodeUsecase) DeleteNode(n *presenter.Node) (*model.Node, error) {
	dn := model.Node{UUID: n.UUID}
	if err := nu.nr.Delete(&dn); err != nil {
		return nil, err
	}
	return &dn, nil
}
