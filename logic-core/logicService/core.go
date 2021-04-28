package logicService

import (
	"errors"
	"fmt"
	"log"

	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
	"github.com/KumKeeHyun/toiot/logic-core/logicService/logic"
)

type logicService struct {
	mux
}

type mux struct {
	chTable map[int]map[int]chan model.LogicData
}

func NewLogicService() *logicService {
	return &logicService{
		mux{
			chTable: make(map[int]map[int]chan model.LogicData),
		},
	}
}

func (m *mux) CreateAndStartLogic(l *model.Logic) error {
	listen := make(chan model.LogicData, 100)
	log.Println("in CreateAndStartLogic, l=", l)
	lchs, ok := m.chTable[l.SensorID]
	if !ok {
		m.chTable[l.SensorID] = make(map[int]chan model.LogicData)
		lchs, _ = m.chTable[l.SensorID]
	}
	if _, ok := lchs[l.ID]; ok {
		close(listen)
		return errors.New("already exist logic evnet")
	}
	lchs[l.ID] = listen

	elems, err := logic.BuildLogic(l)
	log.Println("end BuildLogic")
	if err != nil {
		log.Println("end BuildLogic, return error")
		return err
	}
	go func() {
		log.Println("in CreateAndStartLogic, run go routin")
		for d := range listen {
			log.Println("in CreateAndStartLogic, exec")
			elems.Exec(&d)
		}
	}()

	return nil
}

func (m *mux) RemoveLogic(sid, lid int) error {
	ch, ok := m.chTable[sid][lid]
	if !ok {
		fmt.Errorf("GetLogicChans : cannot find listen channels")
	}
	close(ch)
	delete(m.chTable[sid], lid)
	if len(m.chTable[sid]) == 0 {
		delete(m.chTable, sid)
	}
	return nil
}

func (m *mux) GetLogicChans(sid int) (map[int]chan model.LogicData, error) {
	lchs, ok := m.chTable[sid]
	if !ok || len(lchs) == 0 {
		return nil, fmt.Errorf("GetLogicChans : cannot find listen channels")
	}
	return lchs, nil
}
