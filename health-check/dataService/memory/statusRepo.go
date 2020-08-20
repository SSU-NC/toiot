package memory

import (
	"errors"
	"sync"

	"github.com/KumKeeHyun/PDK/health-check/domain/model"
)

type statusRepo struct {
	mu    *sync.RWMutex
	table map[string]model.Status
}

var statusTable *statusRepo

func NewStatusRepo() *statusRepo {
	if statusTable != nil {
		return statusTable
	}

	statusTable := &statusRepo{
		mu:    &sync.RWMutex{},
		table: map[string]model.Status{},
	}
	return statusTable
}

func (sr *statusRepo) StartAtomic() {
	sr.mu.Lock()
}

func (sr *statusRepo) EndAtomic() {
	sr.mu.Unlock()
}

func (sr *statusRepo) GetKeys() []string {
	keys := make([]string, 0, len(sr.table))
	for k := range sr.table {
		keys = append(keys, k)
	}
	return keys
}

func (sr *statusRepo) Create(key string, value model.Status) error {
	if _, ok := sr.table[key]; ok {
		return errors.New("statusRepo : alreay exist status")
	}
	sr.table[key] = value
	return nil
}

func (sr *statusRepo) Delete(key string) error {
	if _, ok := sr.table[key]; !ok {
		return errors.New("statusRepo : cannot find status")
	}
	delete(sr.table, key)
	return nil
}

func (sr *statusRepo) Get(key string) (model.Status, error) {
	if s, ok := sr.table[key]; !ok {
		return model.Status{}, errors.New("statusRepo : cannot find status")
	} else {
		return s, nil
	}

}

func (sr *statusRepo) Update(key string, value model.Status) error {
	if _, ok := sr.table[key]; !ok {
		return errors.New("statusRepo : cannot find status")
	}

	sr.table[key] = value
	return nil
}
