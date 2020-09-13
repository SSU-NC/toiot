package memory

import (
	"errors"
	"sync"

	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
)

var regist *registRepo

func NewRegistRepo() *registRepo {
	if regist != nil {
		return regist
	}

	regist := &registRepo{
		nodeRepo{
			nmu:   &sync.RWMutex{},
			ninfo: make(map[int]model.Node),
		},
		sensorRepo{
			smu:   &sync.RWMutex{},
			sinfo: make(map[int]model.Sensor),
		},
	}

	return regist
}

type registRepo struct {
	nodeRepo
	sensorRepo
}

type nodeRepo struct {
	nmu   *sync.RWMutex
	ninfo map[int]model.Node
}

func (nr *nodeRepo) FindNode(key int) (*model.Node, error) {
	nr.nmu.RLock()
	defer nr.nmu.RUnlock()

	n, ok := nr.ninfo[key]

	if !ok {
		return nil, errors.New("nodeRepo: cannot find node")
	}
	return &n, nil
}

func (nr *nodeRepo) CreateNode(key int, n *model.Node) error {
	_, ok := nr.ninfo[key]
	if ok {
		return errors.New("nodeRepo: already exist node")
	}
	nr.ninfo[key] = *n
	return nil
}

func (nr *nodeRepo) DeleteNode(key int) error {
	_, ok := nr.ninfo[key]
	if !ok {
		return errors.New("nodeRepo: cannot find node")
	}
	delete(nr.ninfo, key)
	return nil
}

type sensorRepo struct {
	smu   *sync.RWMutex
	sinfo map[int]model.Sensor
}

func (sr *sensorRepo) FindSensor(key int) (*model.Sensor, error) {
	sr.smu.RLock()
	defer sr.smu.RUnlock()

	s, ok := sr.sinfo[key]
	if !ok {
		return nil, errors.New("nodeRepo: cannot find sensor")
	}
	return &s, nil
}

func (sr *sensorRepo) CreateSensor(key int, s *model.Sensor) error {
	_, ok := sr.sinfo[key]
	if ok {
		return errors.New("nodeRepo: already exist sensor")
	}
	sr.sinfo[key] = *s
	return nil
}

func (sr *sensorRepo) DeleteSensor(key int) error {
	_, ok := sr.sinfo[key]
	if !ok {
		return errors.New("nodeRepo: cannot find sensor")
	}
	delete(sr.sinfo, key)
	return nil
}
