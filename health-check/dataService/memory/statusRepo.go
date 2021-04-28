package memory

import (
	"log"
	"sync"
	"time"

	"github.com/KumKeeHyun/toiot/health-check/adapter"
	"github.com/KumKeeHyun/toiot/health-check/domain/model"
)

var (
	loc, _  = time.LoadLocation("Asia/Seoul")
	timeFmt = "2006-01-02 15:04:05"
)

type statusRepo struct {
	mu *sync.RWMutex
	// map[sinkID]map[nodeID]
	table map[int]map[int]model.Status // 이차원 맵 model.Status가 원소 (0,1,2)
}

var statusTable *statusRepo

func NewStatusRepo() *statusRepo {
	if statusTable != nil {
		return statusTable
	}

	statusTable := &statusRepo{
		mu:    &sync.RWMutex{},
		table: map[int]map[int]model.Status{},
	}
	return statusTable
}

func (sr *statusRepo) Lock() {
	sr.mu.Lock()
}

func (sr *statusRepo) Unlock() {
	sr.mu.Unlock()
}

func (sr *statusRepo) UpdateTable(states adapter.States) model.SinkStatus { // ID 번째 싱크를 업데이트 한다.
	t, err := time.ParseInLocation(timeFmt, states.Timestamp, loc)
	if err != nil {
		t = time.Now()
	}

	sr.mu.Lock()
	defer sr.mu.Unlock()

	if _, ok := sr.table[states.State.SinkID]; !ok {
		sr.table[states.State.SinkID] = map[int]model.Status{}
	}
	return sr.updateNodeStatus(states.State.SinkID, states.State.State, t)
}

func (sr *statusRepo) updateNodeStatus(sinkID int, ns []adapter.NodeState, t time.Time) model.SinkStatus { // 어답더 계층의 NodeState상태와 메모리 계층의 statusRepo의 status table을 동기화시켜 주는 것
	res := []model.NodeStatus{}
	rres := model.SinkStatus{}
	nsTable := map[int]bool{}

	// update the status checked from the sink node
	for _, v := range ns { // v는 NodeSate 배열 중 한 원소
		nsTable[v.NodeID] = true
		nodeState, ok := sr.table[sinkID][v.NodeID]
		// if new nodeState, regist new state
		if !ok {
			tempState := model.NewStatus(v.State, t)
			sr.table[sinkID][v.NodeID] = tempState
			res = append(res, model.NodeStatus{NodeID: v.NodeID, State: tempState.State, Battery: v.Battery})
			continue
		}
		if isChanged := nodeState.UpdateState(v.State, t); isChanged {
			res = append(res, model.NodeStatus{NodeID: v.NodeID, State: nodeState.State, Battery: v.Battery})
		}
		sr.table[sinkID][v.NodeID] = nodeState
	}

	// if the state is not confirmed from the sink node
	// check timeout and drop state from table
	// sr.table[sinkID][K]랑 nsTable[k]가 존재하지 않을 경우 제거, 존재할 경우 업데이트
	for k, v := range sr.table[sinkID] {
		if _, ok := nsTable[k]; !ok {
			if v.CheckDrop() {
				delete(sr.table[sinkID], k)
			} else {
				sr.table[sinkID][k] = v
				res = append(res, model.NodeStatus{NodeID: k, State: v.State}) // , Battery: v.Battery
			}

		}
	}
	for i, j := range sr.table {
		log.Println(i, ":", j)
	}
	rres = model.SinkStatus{SinkID: sinkID, Satates: res}
	return rres
}

// func (sr *statusRepo) GetKeys() []string {
// 	keys := make([]string, 0, len(sr.table))
// 	for k := range sr.table {
// 		keys = append(keys, k)
// 	}
// 	return keys
// }

// func (sr *statusRepo) Create(key string, value model.Status) error {
// 	if _, ok := sr.table[key]; ok {
// 		return errors.New("statusRepo : alreay exist status")
// 	}
// 	sr.table[key] = value
// 	return nil
// }

// func (sr *statusRepo) Delete(key string) error {
// 	if _, ok := sr.table[key]; !ok {
// 		return errors.New("statusRepo : cannot find status")
// 	}
// 	delete(sr.table, key)
// 	return nil
// }

// func (sr *statusRepo) Get(key string) (model.Status, error) {
// 	if s, ok := sr.table[key]; !ok {
// 		return model.Status{}, errors.New("statusRepo : cannot find status")
// 	} else {
// 		return s, nil
// 	}

// }

// func (sr *statusRepo) Update(key string, value model.Status) error {
// 	if _, ok := sr.table[key]; !ok {
// 		return errors.New("statusRepo : cannot find status")
// 	}

// 	sr.table[key] = value
// 	return nil
// }

// func (sr *statusRepo) GetHealthInfo() []adapter.HealthInfo {
// 	// sr.mu.RLock()
// 	// defer sr.mu.Unlock()

// 	res := make([]adapter.HealthInfo, 0, len(sr.table))
// 	for k, v := range sr.table {
// 		res = append(res, adapter.HealthInfo{
// 			UUID:  k,
// 			State: v.State,
// 		})
// 	}
// 	return res
// }
