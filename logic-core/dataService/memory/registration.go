package memory

import (
	"errors"
	"log"
	"sync"

	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
)

var regist *registRepo

func NewRegistRepo() (*registRepo, map[int]model.Sink) {
	if regist != nil {
		return regist, nil
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
		sinkAddrRepo{
			samu:  &sync.RWMutex{},
			addrs: make(map[int]model.Sink),
		},
		nodeInfoRepo{
			nmu:   &sync.RWMutex{},
			ninfo: make(map[int]model.Nodeinfo),
		},
	}

	return regist, regist.addrs
}

type registRepo struct {
	nodeRepo
	sensorRepo
	sinkAddrRepo
	nodeInfoRepo
}

type nodeRepo struct {
	nmu   *sync.RWMutex
	ninfo map[int]model.Node
}
type nodeInfoRepo struct {
	nmu   *sync.RWMutex
	ninfo map[int]model.Nodeinfo
}

func (nir *nodeInfoRepo) AppendNodeMap(nid int, sid int) error {
	nir.nmu.RLock()
	defer nir.nmu.RUnlock()

	_, ok := nir.ninfo[nid]

	if ok {
		return errors.New("nodeInfoRepo: already exist nid")
	}
	// var ni model.Nodeinfo
	// ni.SinkID = sid
	ni := model.Nodeinfo{SinkID: sid}

	nir.ninfo[nid] = ni

	log.Println("test >>>>>> in memory/AppendNodeMap, sinkID : ", ni, "sinkADDR : ")
	return nil

}

func (nir *nodeInfoRepo) GetSid(nid int) (*model.Nodeinfo, error) {
	nir.nmu.RLock()
	defer nir.nmu.RUnlock()

	n, ok := nir.ninfo[nid]

	if !ok {
		return nil, errors.New("nodeRepo: cannot find node")
	}
	return &n, nil
}

type sinkAddrRepo struct {
	samu  *sync.RWMutex
	addrs map[int]model.Sink
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

func (sar *sinkAddrRepo) AppendSinkAddr(sid int, s *string) error {
	sar.samu.RLock()
	defer sar.samu.RUnlock()
	_, ok := sar.addrs[sid]
	if ok {
		return errors.New("sinkAddrRepo: already exist sink")
	}
	var sink model.Sink
	sink.Addr = *s
	sar.addrs[sid] = sink
	log.Println("test >>>>>> in memory/appendSinkAddr, sinkID : ", sid, "sinkADDR : ", *s)
	return nil
}

// func (sar *sinkAddrRepo) GetSinkAddrMap() *map[int]model.Sink {
// 	sar.samu.RLock()
// 	defer sar.samu.RUnlock()

// 	return &sar.addrs
// }
