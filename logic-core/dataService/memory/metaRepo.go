package memory

import (
	"errors"
	"sync"

	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
)

var metaData *metaRepo

var nodeInfo = map[string]model.Node{
	"node_1": model.Node{
		Name:  "hn-1floor",
		Group: "soongsil",
	},
	"node_2": model.Node{
		Name:  "hn-2floor",
		Group: "soongsil",
	},
}

var sensorInfo = map[string]model.Sensor{
	"sensor_0": model.Sensor{
		Name: "temporature",
		ValueNames: []string{
			"degree",
		},
	},
	"sensor_1": model.Sensor{
		Name: "dust",
		ValueNames: []string{
			"pm10",
			"pm2.5",
		},
	},
	"sensor_2": model.Sensor{
		Name: "air",
		ValueNames: []string{
			"co",
			"no",
			"so",
		},
	},
}

func NewMetaRepo() *metaRepo {
	if metaData != nil {
		return metaData
	}

	metaData := &metaRepo{
		nodeRepo{
			nmu:   &sync.RWMutex{},
			ninfo: make(map[string]model.Node),
		},
		sensorRepo{
			smu:   &sync.RWMutex{},
			sinfo: make(map[string]model.Sensor),
		},
	}

	metaData.ninfo = nodeInfo
	metaData.sinfo = sensorInfo

	return metaData
}

type metaRepo struct {
	nodeRepo
	sensorRepo
}

type nodeRepo struct {
	nmu   *sync.RWMutex
	ninfo map[string]model.Node
}

func (nr *nodeRepo) GetNode(key string) (*model.Node, error) {
	nr.nmu.RLock()
	defer nr.nmu.RUnlock()

	n, ok := nr.ninfo[key]
	if !ok {
		return nil, errors.New("nodeRepo: cannot find node")
	}
	return &n, nil
}

func (nr *nodeRepo) NewNode(key string, n *model.Node) error {
	nr.nmu.Lock()
	defer nr.nmu.Unlock()

	_, ok := nr.ninfo[key]
	if ok {
		return errors.New("nodeRepo: already exist node")
	}
	nr.ninfo[key] = *n
	return nil
}

func (nr *nodeRepo) DelNode(key string) error {
	nr.nmu.Lock()
	defer nr.nmu.Unlock()

	_, ok := nr.ninfo[key]
	if !ok {
		return errors.New("nodeRepo: cannot delete node")
	}
	delete(nr.ninfo, key)
	return nil
}

type sensorRepo struct {
	smu   *sync.RWMutex
	sinfo map[string]model.Sensor
}

func (sr *sensorRepo) GetSensor(key string) (*model.Sensor, error) {
	sr.smu.RLock()
	defer sr.smu.RUnlock()

	s, ok := sr.sinfo[key]
	if !ok {
		return nil, errors.New("nodeRepo: cannot find sensor")
	}
	return &s, nil
}

func (sr *sensorRepo) NewSensor(key string, s *model.Sensor) error {
	sr.smu.Lock()
	defer sr.smu.Unlock()

	_, ok := sr.sinfo[key]
	if ok {
		return errors.New("nodeRepo: already exist sensor")
	}
	sr.sinfo[key] = *s
	return nil
}

func (sr *sensorRepo) DelSensor(key string) error {
	sr.smu.Lock()
	defer sr.smu.Unlock()

	_, ok := sr.sinfo[key]
	if !ok {
		return errors.New("nodeRepo: cannot delete sensor")
	}
	delete(sr.sinfo, key)
	return nil
}
