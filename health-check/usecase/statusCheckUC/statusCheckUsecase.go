package statusCheckUC

import (
	"time"

	"github.com/KumKeeHyun/PDK/health-check/domain/repository"
	"github.com/KumKeeHyun/PDK/health-check/setting"
)

type statusCheckUsecase struct {
	sr    repository.StatusRepo
	event chan struct{}
}

func NewStatusCheckUsecase(sr repository.StatusRepo, e chan struct{}) *statusCheckUsecase {
	su := &statusCheckUsecase{
		sr:    sr,
		event: e,
	}

	go func() {
		tick := time.Tick(time.Duration(setting.StatusSetting.Tick) * time.Second)
		for {
			select {
			case <-tick:
				su.check()
			}
		}
	}()

	return su
}

func (su *statusCheckUsecase) check() {
	su.sr.StartAtomic()
	defer su.sr.EndAtomic()

	keys := su.sr.GetKeys()
	change := false
	for _, k := range keys {
		if s, err := su.sr.Get(k); err != nil {
			continue
		} else {
			if s.CheckDrop() {
				if err := su.sr.Delete(k); err != nil {
					// TODO
				}
				continue
			}
			change = s.CheckCnt()
			if err := su.sr.Update(k, s); err != nil {
				// TODO
			}
		}
	}
	if change {
		su.event <- struct{}{}
	}
}
