package kafkaConsumer

import (
	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
	"github.com/KumKeeHyun/PDK/logic-core/setting"
)

var kafkaConsumer *group

type group struct {
	cs  []*consumer
	out chan model.KafkaData
}

func NewKafkaConsumer() *group {
	if kafkaConsumer != nil {
		return kafkaConsumer
	}

	outBufSize := 100
	numOfConsumers := setting.KafkaSetting.NumOfConsumers

	kafkaConsumer = &group{
		cs:  make([]*consumer, numOfConsumers),
		out: make(chan model.KafkaData, outBufSize),
	}

	for i := 0; i < numOfConsumers; i++ {
		kafkaConsumer.cs[i] = NewConsumer()
		if kafkaConsumer.cs[i] != nil {
			go kafkaConsumer.cs[i].run(kafkaConsumer.out)
		}
	}
	return kafkaConsumer
}

func (g *group) GetOutput() <-chan model.KafkaData {
	if g != nil {
		return g.out
	}
	return nil
}
