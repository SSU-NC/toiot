package kafkaPipe

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/KumKeeHyun/PDK/logic-core/setting"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type ConsumerGroup struct {
	Group []*kafka.Consumer
}

var consumerGroup *ConsumerGroup

const BUFSIZE = 100

func Setup() error {
	config := &kafka.ConfigMap{
		"bootstrap.servers":               setting.KafkaSetting.Broker,
		"group.id":                        setting.KafkaSetting.GroupID,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"enable.partition.eof":            true,
		"auto.offset.reset":               "earliest",
	}

	consumerGroup = &ConsumerGroup{
		Group: make([]*kafka.Consumer, 0),
	}

	for i := 0; i < setting.KafkaSetting.NumOfConsumers; i++ {
		nc, err := kafka.NewConsumer(config)
		if err != nil {
			return err
		}
		if err = nc.SubscribeTopics(setting.KafkaSetting.Topics, nil); err != nil {
			return err
		}
		consumerGroup.Group = append(consumerGroup.Group, nc)
	}
	return nil
}

func StartConsumer() <-chan KafkaData {
	out := make(chan KafkaData, BUFSIZE)

	for i := 0; i < setting.KafkaSetting.NumOfConsumers; i++ {
		go ConsumKafka(consumerGroup.Group[i], i, out)
	}

	return out
}

func ConsumKafka(c *kafka.Consumer, i int, out chan<- KafkaData) {
	for ev := range c.Events() {
		switch e := ev.(type) {
		case kafka.AssignedPartitions:
			fmt.Fprintf(os.Stderr, "%% %v\n", e)
			c.Assign(e.Partitions)
		case kafka.RevokedPartitions:
			fmt.Fprintf(os.Stderr, "%% %v\n", e)
			c.Unassign()
		case *kafka.Message:
			fmt.Fprintf(os.Stderr, "Get msg from %d consumer\n", i)
			data := KafkaData{
				Key: string(e.Key),
			}
			if err := json.Unmarshal(e.Value, &data.Value); err != nil {
				continue
			}
			out <- data
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
		}
	}
}
