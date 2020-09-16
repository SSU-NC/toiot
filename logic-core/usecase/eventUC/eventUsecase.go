package eventUC

import (
	"github.com/KumKeeHyun/toiot/logic-core/adapter"
	"github.com/KumKeeHyun/toiot/logic-core/domain/repository"
	"github.com/KumKeeHyun/toiot/logic-core/domain/service"
)

type eventUsecase struct {
	rr repository.RegistRepo
	ls service.LogicService
}

func NewEventUsecase(rr repository.RegistRepo, ls service.LogicService) *eventUsecase {
	return &eventUsecase{
		rr: rr,
		ls: ls,
	}
}

func (eu *eventUsecase) DeleteSink(nl []adapter.Node) error {
	for _, n := range nl {
		eu.rr.DeleteNode(n.ID)
	}
	return nil
}

func (eu *eventUsecase) CreateNode(n *adapter.Node, sn string) error {
	mn, asl := adapter.NodeToModel(n, sn)
	eu.rr.CreateNode(n.ID, &mn)

	all := []adapter.Logic{}
	for _, as := range asl {
		ms, tempAll := adapter.SensorToModel(&as)
		all = append(all, tempAll...)
		eu.rr.CreateSensor(as.ID, &ms)
	}
	mll := adapter.LogicsToModels(all)
	for _, ml := range mll {
		eu.ls.CreateAndStartLogic(&ml)
	}

	return nil
}

func (eu *eventUsecase) DeleteNode(n *adapter.Node) error {
	return eu.rr.DeleteNode(n.ID)
}

func (eu *eventUsecase) DeleteSensor(s *adapter.Sensor) error {
	for _, l := range s.Logics {
		eu.ls.RemoveLogic(l.SensorID, l.ID)
	}
	return eu.rr.DeleteSensor(s.ID)
}

func (eu *eventUsecase) CreateLogic(l *adapter.Logic) error {
	if ml, err := adapter.LogicToModel(l); err != nil {
		return err
	} else {
		return eu.ls.CreateAndStartLogic(&ml)
	}
}

func (eu *eventUsecase) DeleteLogic(l *adapter.Logic) error {
	return eu.ls.RemoveLogic(l.SensorID, l.ID)
}
