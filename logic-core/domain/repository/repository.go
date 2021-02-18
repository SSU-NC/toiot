package repository

import (
	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
)

type RegistRepo interface {
	FindNode(key int) (*model.Node, error)
	CreateNode(key int, n *model.Node) error
	DeleteNode(key int) error
	// AppendNodeInfo(nid int, sid int) error
	// GetSid(nid int) (*model.Nodeinfo, error)
	FindSensor(key int) (*model.Sensor, error)
	CreateSensor(key int, s *model.Sensor) error
	DeleteSensor(key int) error
	AppendSinkAddr(sid int, s *string) error

	// GetSinkAddrMap() *map[int]model.Sink
}
