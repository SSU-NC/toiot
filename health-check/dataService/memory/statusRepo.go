package memory

import (
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
	table map[int]map[int]model.Status
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

func (sr *statusRepo) UpdateTable(sinkID int, states adapter.States) []model.NodeStatus {
	t, err := time.ParseInLocation(timeFmt, states.Timestamp, loc)
	if err != nil {
		t = time.Now()
	}

	sr.mu.Lock()
	defer sr.mu.Unlock()

	if _, ok := sr.table[sinkID]; !ok {
		sr.table[sinkID] = map[int]model.Status{}
	}
	return sr.updateNodeStatus(sinkID, states.State, t)
}

func (sr *statusRepo) updateNodeStatus(sinkID int, ns []adapter.NodeState, t time.Time) []model.NodeStatus {
	res := []model.NodeStatus{}
	nsTable := map[int]bool{}

	// update the status checked from the sink node
	for _, v := range ns {
		nsTable[v.NodeID] = true
		nodeState, ok := sr.table[sinkID][v.NodeID]
		// if new nodeState, regist new state
		if !ok {
			tempState := model.NewStatus(v.State, t)
			sr.table[sinkID][v.NodeID] = tempState
			res = append(res, model.NodeStatus{NodeID: v.NodeID, State: tempState.State})
			continue
		}
		if isChanged := nodeState.UpdateState(v.State, t); isChanged {
			res = append(res, model.NodeStatus{NodeID: v.NodeID, State: nodeState.State})
		}
		sr.table[sinkID][v.NodeID] = nodeState
	}

	// if the state is not confirmed from the sink node
	// check timeout and drop state from table
	for k, v := range sr.table[sinkID] {
		if _, ok := nsTable[k]; !ok {
			if v.CheckDrop() {
				delete(sr.table[sinkID], k)
			} else {
				sr.table[sinkID][k] = v
				res = append(res, model.NodeStatus{NodeID: k, State: v.State})
			}

		}
	}
	return res
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
