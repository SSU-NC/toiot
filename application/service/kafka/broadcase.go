package kafka

import "github.com/dustin/go-broadcast"

type kafkaManager struct {
	broadcast.Broadcaster
}

var MessageManager *kafkaManager

func Setup() {
	MessageManager = &kafkaManager{
		broadcast.NewBroadcaster(10),
	}
}
