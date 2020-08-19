package usecase

import "github.com/seheee/PDK/logic-core/domain/model"

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
	NewNode(key string, n *model.Node) (*model.Node, error)
	NewSensor(key string, s *model.Sensor) (*model.Sensor, error)
	DeleteNode(key string) error
	DeleteSensor(key string) error
}
