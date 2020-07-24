package model

import (
	"errors"
	"sync"
)

var RegisterRepo = RegisterInfo{
	mu:         &sync.RWMutex{},
	NodeInfo:   map[string]Node{},
	SensorInfo: map[string]Sensor{},
}

type RegisterInfo struct {
	mu         *sync.RWMutex
	NodeInfo   map[string]Node
	SensorInfo map[string]Sensor
}

type InitRegister struct {
	NodeInfo   []AppNode   `json:"node_info"`
	SensorInfo []AppSensor `json:"sensor_info"`
}

func (repo *RegisterInfo) GetNode(k string) (Node, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if n, ok := repo.NodeInfo[k]; ok {
		return n, nil
	} else {
		return n, errors.New("not exist")
	}
}

func (repo *RegisterInfo) AddNode(k string, n Node) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.NodeInfo[k] = n
}

func (repo *RegisterInfo) DelNode(k string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	delete(repo.NodeInfo, k)
}

func (repo *RegisterInfo) GetSensor(k string) (Sensor, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if s, ok := repo.SensorInfo[k]; ok {
		return s, nil
	} else {
		return s, errors.New("not exist")
	}
}

func (repo *RegisterInfo) AddSensor(k string, s Sensor) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.SensorInfo[k] = s
}

func (repo *RegisterInfo) DelSensor(k string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	delete(repo.SensorInfo, k)
}
