package kafkaConsumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/KumKeeHyun/PDK/health-check/adapter.go"
	"github.com/KumKeeHyun/PDK/health-check/setting"
	"github.com/Shopify/sarama"
)

var kafkaConsumer *group

type group struct {
	client sarama.ConsumerGroup
	out    chan adapter.States
}

func NewKafkaConsumer() *group {
	var err error

	if kafkaConsumer != nil {
		return kafkaConsumer
	}

	outBufSize := setting.KafkaSetting.ChanBufSize

	kafkaConsumer = &group{
		out: make(chan adapter.States, outBufSize),
	}

	cfg := sarama.NewConfig()
	cfg.Version = sarama.V0_10_2_0
	cfg.Consumer.Offsets.Initial = sarama.OffsetNewest

	kafkaConsumer.client, err = sarama.NewConsumerGroup([]string{setting.KafkaSetting.Broker}, setting.KafkaSetting.GroupID, cfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	consumer := consumer{
		out:   kafkaConsumer.out,
		ready: make(chan bool),
	}
	go func() {
		for {
			err = kafkaConsumer.client.Consume(ctx, setting.KafkaSetting.Topics, &consumer)
			if err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	return kafkaConsumer
}

func (g *group) GetOutput() <-chan adapter.States {
	if g != nil {
		return g.out
	}
	return nil
}

type consumer struct {
	out   chan adapter.States
	ready chan bool
}

func (consumer *consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		//fmt.Printf("sarama\nValue : %s\n", string(message.Value))
		d := adapter.States{}
		if err := json.Unmarshal(message.Value, &d); err != nil {
			continue
		}
		// TODO : check valid
		consumer.out <- d
	}

	return nil
}
