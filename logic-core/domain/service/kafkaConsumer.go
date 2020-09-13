package service

import "github.com/KumKeeHyun/toiot/logic-core/domain/model"

type KafkaConsumerGroup interface {
	GetOutput() <-chan model.KafkaData

	// IncreaseConsumer() error
	// DecreaseConsumer() error
}
