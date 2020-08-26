package service

import "github.com/seheee/PDK/logic-core/domain/model"

type KafkaConsumerGroup interface {
	GetOutput() <-chan model.KafkaData

	// IncreaseConsumer() error
	// DecreaseConsumer() error
}
