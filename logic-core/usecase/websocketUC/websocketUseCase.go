package websocketUC

import (
	"fmt"
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
		for ev := range wu.event {
			fmt.Println("broadcast\n", ev)
			wu.Submit(ev)
		}
	}()

	return wu
}