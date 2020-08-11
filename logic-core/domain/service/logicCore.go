package service

import "github.com/seheee/PDK/logic-core/domain/model"

type LogicCore interface {
	CreateAndStartLogic(r *model.RingRequest)
	GetLogicChans(key string) map[string]chan model.LogicData
	RemoveLogic(lname string) error
	RemoveLogicsBySID(sid string) error
}
