package websocketUC

import (
	"github.com/dustin/go-broadcast"
)

type websocketUsecase struct {
	event chan interface{}
	broadcast.Broadcaster
}

func NewWebsocketUsecase(e chan interface{}) *websocketUsecase {
	wu := &websocketUsecase{
		event:       e,
		Broadcaster: broadcast.NewBroadcaster(10),
	}

	go func() {
		for msg := range wu.event {
			wu.Submit(msg)
		}
	}()

	return wu
}
