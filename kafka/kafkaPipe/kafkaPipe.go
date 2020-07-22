package kafkaPipe

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/KumKeeHyun/PDK/kafka/setting"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var consumer *kafka.Consumer

const BUFSIZE = 1

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
	c, err := kafka.NewConsumer(config)
	if err != nil {
		return err
	}
	consumer = c
	err = consumer.SubscribeTopics(setting.KafkaSetting.Topics, nil)
	return err
}

func ConsumKafka(end chan os.Signal) <-chan KafkaData {
	out := make(chan KafkaData, BUFSIZE)
	run := true
	go func() {
		defer func() {
			consumer.Close()
			close(out)
		}()
	end:
		for run == true {
			select {
			case <-end:
				run = false
				fmt.Println("end pipe")
				break end
			case ev := <-consumer.Events():
				switch e := ev.(type) {
				case kafka.AssignedPartitions:
					fmt.Fprintf(os.Stderr, "%% %v\n", e)
					consumer.Assign(e.Partitions)
				case kafka.RevokedPartitions:
					fmt.Fprintf(os.Stderr, "%% %v\n", e)
					consumer.Unassign()
				case *kafka.Message:
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

	}()
	return out
}
