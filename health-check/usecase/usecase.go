package usecase

import (
	"github.com/dustin/go-broadcast"
)

type HealthCheckkUsecase interface {
	healthCheck()
}

type WebsocketUsecase interface {
	broadcast.Broadcaster
}
