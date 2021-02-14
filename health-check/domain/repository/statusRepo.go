package repository

import (
	"github.com/KumKeeHyun/toiot/health-check/adapter"
	"github.com/KumKeeHyun/toiot/health-check/domain/model"
)

type StatusRepo interface {
	UpdateTable(states adapter.States) model.SinkStatus
}
