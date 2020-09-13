package eventUsecase

import (
	"fmt"
	"sync"
	"time"

	"github.com/KumKeeHyun/toiot/application/domain/model"
	"github.com/KumKeeHyun/toiot/application/domain/repository"
)

type eventUsecase struct {
	requestRetry []pingRequest
	lsr          repository.LogicServiceRepo
}

func NewEventUsecase(lsr repository.LogicServiceRepo) *eventUsecase {
	eu := &eventUsecase{
		requestRetry: []pingRequest{},
		lsr:          lsr,
	}
	tick := time.Tick(10 * time.Second)
	go func() {
		for {
			select {
			case <-tick:
				eu.CheckAndUnregistLogicServices()
			}
		}
	}()
	return eu
}

// `{
// 	"addr" : "localhost:8082",
// 	"topic" : {
// 		"name":"sensors"
// 	}
// }`
func (eu *eventUsecase) RegistLogicService(l *model.LogicService) error {
	return eu.lsr.Create(l)
}

func (eu *eventUsecase) CheckAndUnregistLogicServices() error {
	var wg sync.WaitGroup
	for _, pr := range eu.requestRetry {
		wg.Add(1)
		go func(_pr pingRequest) {
			if err := _pr.ping(); err != nil {
				fmt.Println(err.Error())
				eu.lsr.Delete(&_pr.ls)
			}
			wg.Done()
		}(pr)
	}
	wg.Wait()
	eu.requestRetry = []pingRequest{}

	return nil
}

type EVENT int

const (
	DeleteSink EVENT = iota
	CreateNode
	DeleteNode
	DeleteSensor
	CreateLogic
	DeleteLogic
)

var EventPath = [...]string{
	"/event/sink/delete",
	"/event/node/create",
	"/event/node/delete",
	"/event/sensor/delete",
	"/event/logic/create",
	"/event/logic/delete",
}

type pingRequest struct {
	ls   model.LogicService
	e    EVENT
	body interface{}
}

func (pr *pingRequest) ping() error {
	url := makeUrl(pr.ls.Addr, EventPath[pr.e])

	resp, _ := pingClient.R().SetBody(pr.body).Post(url)
	if resp.IsSuccess() {
		return nil
	}
	return fmt.Errorf("ping fail : %v", *pr)
}
