package eventUsecase

import (
	"fmt"
	"sync"
	"time"

	"github.com/KumKeeHyun/toiot/application/domain/model"
	"github.com/KumKeeHyun/toiot/application/domain/repository"
)

type eventUsecase struct {
	lsr repository.LogicServiceRepo
}

func NewEventUsecase(lsr repository.LogicServiceRepo) *eventUsecase {
	eu := &eventUsecase{
		lsr: lsr,
	}
	tick := time.Tick(5 * time.Second)
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
	ls, err := eu.lsr.FindsWithTopic()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, l := range ls {
		wg.Add(1)
		go func() {
			if err := ping(l); err != nil {
				fmt.Println(err.Error())
				eu.lsr.Delete(&l)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return nil
}
