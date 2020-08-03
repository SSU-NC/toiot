package kafkaConsumer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
	"github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type consumer struct {
	c   *kafka.Consumer
	ctx chan struct{}
}

func NewConsumer() *consumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers":               setting.KafkaSetting.Broker,
		"group.id":                        setting.KafkaSetting.GroupID,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"enable.partition.eof":            true,
		"auto.offset.reset":               "earliest",
	}

	c, err := kafka.NewConsumer(config)
	if err != nil {
		return nil
	}

	if err = c.SubscribeTopics(setting.KafkaSetting.Topics, nil); err != nil {
		return nil
	}

	return &consumer{
		c:   c,
		ctx: make(chan struct{}, 1),
	}
}

func (c *consumer) run(out chan<- model.KafkaData) {
	for {
		select {
		case <-c.ctx:
			fmt.Printf("consumer stop\n")
			c.c.Close()
			return
		case ev := <-c.c.Events():
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("key : %s, Value : %s\n", string(e.Key), string(e.Value))
				d := model.KafkaData{
					Key: string(e.Key),
				}
				if err := json.Unmarshal(e.Value, &d.Value); err != nil {
					continue
				}
				out <- d
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.c.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.c.Unassign()
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}
}

func (c *consumer) stop() {
	c.ctx <- struct{}{}
}
