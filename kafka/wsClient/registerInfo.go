package wsClient

import (
	"encoding/json"
	"fmt"
	"sync"
)

type registerRepository struct {
	Mu   *sync.RWMutex
	Info RegisterInfo
}

var Repo *registerRepository

func NewRegisterRepo() *registerRepository {
	return &registerRepository{
		Mu: &sync.RWMutex{},
		Info: RegisterInfo{
			NodeInfo:   map[string]Node{},
			SensorInfo: map[string]Sensor{},
		},
	}
}

func InitMessage(msg []byte) {
	Repo.Mu.Lock()
	defer Repo.Mu.Unlock()

	var initMsg RegisterInit
	if err := json.Unmarshal(msg, &initMsg); err != nil {
		fmt.Printf("json unmarshal fail : %s\n", err.Error())
		return
	}

	for _, n := range initMsg.NodeInfo {
		Repo.Info.NodeInfo[n.UUID] = n
	}
	for _, s := range initMsg.SensorInfo {
		Repo.Info.SensorInfo[s.UUID] = s
	}
}

func NewNodeMessage(msg []byte) {
	Repo.Mu.Lock()
	defer Repo.Mu.Unlock()

	var node Node
	if err := json.Unmarshal(msg, &node); err != nil {
		fmt.Printf("json unmarshal fail : %s\n", err.Error())
		return
	}
	Repo.Info.NodeInfo[node.UUID] = node
}

func DeleteNodeMessage(msg []byte) {
	Repo.Mu.Lock()
	defer Repo.Mu.Unlock()

	var node Node
	if err := json.Unmarshal(msg, &node); err != nil {
		fmt.Printf("json unmarshal fail : %s\n", err.Error())
		return
	}

	delete(Repo.Info.NodeInfo, node.UUID)
}

func NewSensorMessage(msg []byte) {
	Repo.Mu.Lock()
	defer Repo.Mu.Unlock()

	var sensor Sensor
	if err := json.Unmarshal(msg, &sensor); err != nil {
		fmt.Printf("json unmarshal fail : %s\n", err.Error())
		return
	}
	Repo.Info.SensorInfo[sensor.UUID] = sensor
}

func DeleteSensorMessage(msg []byte) {
	Repo.Mu.Lock()
	defer Repo.Mu.Unlock()

	var sensor Sensor
	if err := json.Unmarshal(msg, &sensor); err != nil {
		fmt.Printf("json unmarshal fail : %s\n", err.Error())
		return
	}
	delete(Repo.Info.SensorInfo, sensor.UUID)
}
