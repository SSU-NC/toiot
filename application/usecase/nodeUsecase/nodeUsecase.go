package nodeUsecase

import (
	"github.com/seheee/PDK/application/domain/model"
	"github.com/seheee/PDK/application/domain/repository"
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

func (nu *nodeUsecase) GetAllNodes() ([]model.Node, error) {
	ns, err := nu.nr.GetAll()
	if err != nil {
		return nil, err
	}
	return ns, nil
}

func (nu *nodeUsecase) GetAllNodesWithSensors() ([]model.Node, error) {
	ns, err := nu.nr.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range ns {
		ns[i].Sensors, err = nu.sr.GetByNodeUUID(ns[i].UUID)
		if err != nil {
			return nil, err
		}
	}
	return ns, nil
}

func (nu *nodeUsecase) GetAllNodesWithSensorsWithValues() ([]model.Node, error) {
	ns, err := nu.nr.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range ns {
		ns[i].Sensors, err = nu.sr.GetByNodeUUIDWithValues(ns[i].UUID)
		if err != nil {
			return nil, err
		}
	}
	return ns, nil
}

func (nu *nodeUsecase) GetNodesByUUID(ids []string) ([]model.Node, error) {
	ns, err := nu.nr.GetByUUIDs(ids)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

func (nu *nodeUsecase) GetNodeByUUID(uuid string) (*model.Node, error) {
	n, err := nu.nr.GetByUUID(uuid)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (nu *nodeUsecase) GetNodeByUUIDWithSensors(uuid string) (*model.Node, error) {
	n, err := nu.nr.GetByUUID(uuid)
	if err != nil {
		return nil, err
	}
	if n.Sensors, err = nu.sr.GetByNodeUUIDWithValues(n.UUID); err != nil {
		return nil, err
	}
	return n, nil
}

func (nu *nodeUsecase) GetNodesBySinkID(sinkID uint) ([]model.Node, error) {
	ns, err := nu.nr.GetBySinkID(sinkID)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

func (nu *nodeUsecase) RegisterNode(n *model.Node) (*model.Node, error) {
	newNode := model.NewNode(n.Name, n.Group, n.LocLat, n.LocLon, n.SinkID)
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

func (nu *nodeUsecase) DeleteNode(n *model.Node) (*model.Node, error) {
	dn := model.Node{UUID: n.UUID}
	if err := nu.nr.Delete(&dn); err != nil {
		return nil, err
	}
	return &dn, nil
}
