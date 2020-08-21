package websocketUC

import (
	"github.com/dustin/go-broadcast"
)

type websocketUsecase struct {
	event chan struct{}
	broadcast.Broadcaster
}

func NewWebsocketUsecase(e chan struct{}) *websocketUsecase {
	wu := &websocketUsecase{
		Broadcaster: broadcast.NewBroadcaster(10),
	}

	go func() {
		for _ = range wu.event {
			his := wu.sr.GetHealthInfo()
			fmt.Println("broadcast\n", his)
			wu.Submit(his)
		}
	}()

	return wu
}