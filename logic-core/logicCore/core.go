package logicCore

import (
	"errors"

	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
)

type logicCore struct {
	mux
}

type mux struct {
	chTable    map[string]map[string]chan model.LogicData
	logicTable map[string]string
}

func NewLogicCore() *logicCore {
	return &logicCore{
		mux{
			chTable:    make(map[string]map[string]chan model.LogicData),
			logicTable: make(map[string]string),
		},
	}
}

func (m *mux) CreateAndStartLogic(r *model.ChainRequest) {
	listen := make(chan model.LogicData, 100)
	lchs, ok := m.chTable[r.SID]
	if !ok {
		m.chTable[r.SID] = make(map[string]chan model.LogicData)
		lchs, _ = m.chTable[r.SID]
	}
	lchs[r.Name] = listen

	chain := chainFactory(r.Rings)
	for d := range listen {
		chain.execute(&d)
	}
}

func (m *mux) GetLogicChans(key string) map[string]chan model.LogicData {
	lchs, ok := m.chTable[key]
	if !ok {
		return nil
	}
	return lchs
}

func (m *mux) RemoveLogic(lname string) error {
	sid, ok := m.logicTable[lname]
	if !ok {
		errors.New("cannot find logicChain " + lname)
	}
	ch, _ := m.chTable[sid][lname]
	close(ch)
	return nil
}

func (m *mux) RemoveLogicsBySID(sid string) error {
	lchs, ok := m.chTable[sid]
	if !ok {
		errors.New("there is no sensor " + sid)
	}
	for _, ch := range lchs {
		close(ch)
	}
	return nil
}
