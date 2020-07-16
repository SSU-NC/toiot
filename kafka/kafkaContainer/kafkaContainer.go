package kafkaContainer

import (
	"fmt"
	"os"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func MakeConfig(broker, group string) *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers":               broker,
		"group.id":                        group,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		// Enable generation of PartitionEOF when the
		// end of a partition is reached.
		"enable.partition.eof": true,
		"auto.offset.reset":    "earliest",
	}
}

func RunKafkaConsumer(c *kafka.Consumer, output chan string, e chan int) {
	run := true
	go func() {
		for run == true {
			select {
			case <-e:
				fmt.Printf("exit kafkaconsumer\n")
				run = false

			case ev := <-c.Events():
				switch e := ev.(type) {
				case kafka.AssignedPartitions:
					fmt.Fprintf(os.Stderr, "%% %v\n", e)
					c.Assign(e.Partitions)
				case kafka.RevokedPartitions:
					fmt.Fprintf(os.Stderr, "%% %v\n", e)
					c.Unassign()
				case *kafka.Message:
					fmt.Printf("%% Message on %s:\n%s\n",
						e.TopicPartition, string(e.Value))
					output <- string(e.Value)
				case kafka.PartitionEOF:
					fmt.Printf("%% Reached %v\n", e)
				case kafka.Error:
					// Errors should generally be considered as informational, the client will try to automatically recover
					fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				}
			}
		}
		c.Close()
	}()
}
