package websocketUC

import (
	"github.com/KumKeeHyun/PDK/health-check/domain/repository"
	"github.com/dustin/go-broadcast"
)

type websocketUsecase struct {
	sr    repository.StatusRepo
	event chan struct{}
	broadcast.Broadcaster
}

func NewWebsocketUsecase(sr repository.StatusRepo, e chan struct{}) *websocketUsecase {
	wu := &websocketUsecase{
		sr:          sr,
		event:       e,
		Broadcaster: broadcast.NewBroadcaster(10),
	}

	go func() {
		for _ = range wu.event {
			his := wu.sr.GetHealthInfo()
			wu.Submit(his)
		}
	}()

	return wu
}
