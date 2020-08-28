package usecase

import (
	"github.com/KumKeeHyun/PDK/logic-core/adapter"
	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
	"github.com/dustin/go-broadcast"
)

// type ConsumerManageUsecase interface {
// 	SetNumOfConsumers(int) (int, error)
// }

type LogicCoreUsecase interface {
	SetLogicChain(r *model.RingRequest) error
	RemoveLogicChain(lname string) error
	RemoveLogicChainsBySID(sid string) error
	GetAllLogics() ([]model.Ring, error)
}

type MetaDataUsecase interface {
	SetMetaInfo(mi adapter.MetaInfo)
	NewNode(key string, n *model.Node) (*model.Node, error)
	NewSensor(key string, s *model.Sensor) (*model.Sensor, error)
	DeleteNode(key string) error
	DeleteSensor(key string) error
}

type WebsocketUsecase interface {
	broadcast.Broadcaster
}
