package healthCheckUC

import (
	"fmt"
	"sync"
	"time"

	"github.com/KumKeeHyun/toiot/health-check/adapter"
	"github.com/KumKeeHyun/toiot/health-check/domain/repository"
	"github.com/KumKeeHyun/toiot/health-check/setting"
	"github.com/go-resty/resty/v2"
)

type healthCheckUsecase struct {
	sr    repository.StatusRepo
	event chan interface{}
}

func NewHealthCheckUsecase(sr repository.StatusRepo, e chan interface{}) *healthCheckUsecase {
	hu := &healthCheckUsecase{
		sr:    sr,
		event: e,
	}

	go func() {
		tick := time.Tick(time.Duration(setting.StatusSetting.Tick) * time.Second)
		for {
			select {
			case <-tick:
				hu.healthCheck()
			}
		}
	}()

	return hu
}

func (hu *healthCheckUsecase) healthCheck() {
	sinks, err := getSinkList()
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	for _, sink := range sinks {
		wg.Add(1)
		go func(s adapter.Sink) {
			res := adapter.States{}
			client := resty.New()
			client.SetTimeout(500 * time.Millisecond)
			resp, _ := client.R().SetResult(&res).Get(fmt.Sprintf("http://%s/health-check", s.Addr))

			if resp.IsSuccess() {
				hu.event <- hu.sr.UpdateTable(s.ID, res)
			}
			wg.Done()
		}(sink)
	}
}
