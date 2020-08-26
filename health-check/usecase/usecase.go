package usecase

import (
	"github.com/seheee/PDK/health-check/adapter.go"
	"github.com/seheee/PDK/health-check/domain/model"
	"github.com/dustin/go-broadcast"
)

type StatusCheckUsecase interface {
	check()
}

type EventUsecase interface {
	GetNodeStatus(ns adapter.NodeState, t string) (model.Status, error)
}

type WebsocketUsecase interface {
	broadcast.Broadcaster
}
