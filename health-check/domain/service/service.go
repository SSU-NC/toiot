package service

import "github.com/KumKeeHyun/PDK/health-check/adapter"

type KafkaConsumer interface {
	GetOutput() <-chan adapter.States
}
