package usecase

import (
	"github.com/KumKeeHyun/toiot/logic-core/adapter"
)

type EventUsecase interface {
	DeleteSink(nl []adapter.Node) error
	CreateNode(n *adapter.Node, sn string) error
	DeleteNode(n *adapter.Node) error
	DeleteSensor(s *adapter.Sensor) error
	CreateLogic(l *adapter.Logic) error
	DeleteLogic(l *adapter.Logic) error
}

type LogicCoreUsecase interface {
	AppendSinkAddr(sa *adapter.SinkAddr) error
}

// type sinkAddrRepo struct {
// 	samu  *sync.RWMutex
// 	addrs []model.Sink
// }

// func (sar *sinkAddrRepo) appendSinkAddr(s model.Sink) error {
// 	sar.addrs = append(sar.addrs, s)
// 	return nil
// }
