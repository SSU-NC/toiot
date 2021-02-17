package service

import "github.com/KumKeeHyun/toiot/logic-core/domain/model"

type LogicService interface {
	CreateAndStartLogic(l *model.Logic) error
	RemoveLogic(sid, lid int) error
	GetLogicChans(sid int) (map[int]chan model.LogicData, error)
	
}
