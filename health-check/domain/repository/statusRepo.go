package repository

import (
	"github.com/KumKeeHyun/PDK/health-check/adapter.go"
	"github.com/KumKeeHyun/PDK/health-check/domain/model"
)

type StatusRepo interface {
	StartAtomic()
	EndAtomic()
	GetKeys() []string
	Create(key string, value model.Status) error
	Delete(key string) error
	Get(key string) (model.Status, error)
	Update(key string, value model.Status) error
	GetHealthInfo() []adapter.HealthInfo
}
