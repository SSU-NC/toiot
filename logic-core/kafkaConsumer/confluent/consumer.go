// /*
// TODO : change kafka module confluent to sarama
// */
package confluent

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	"github.com/KumKeeHyun/PDK/logic-core/adapter"
// 	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
// 	"github.com/KumKeeHyun/PDK/logic-core/setting"
// 	"github.com/confluentinc/confluent-kafka-go/kafka"
// )

// type consumer struct {
// 	c   *kafka.Consumer
// 	ctx chan struct{}
// }

// func NewConsumer() *consumer {
// 	config := &kafka.ConfigMap{
// 		"bootstrap.servers":               setting.Kafkasetting.Broker,
// 		"group.id":                        setting.Kafkasetting.GroupID,
// 		"session.timeout.ms":              6000,
// 		"go.events.channel.enable":        true,
// 		"go.application.rebalance.enable": true,
// 		"enable.partition.eof":            true,
// 		"auto.offset.reset":               "earliest",
// 	}

// 	c, err := kafka.NewConsumer(config)
// 	if err != nil {
// 		return nil
// 	}

// 	if err = c.SubscribeTopics(setting.Kafkasetting.Topics, nil); err != nil {
// 		return nil
// 	}

// 	return &consumer{
// 		c:   c,
// 		ctx: make(chan struct{}, 1),
// 	}
// }

// func (c *consumer) run(out chan<- model.KafkaData) {
// 	for {
// 		select {
// 		case <-c.ctx:
// 			fmt.Printf("consumer stop\n")
// 			c.c.Close()
// 			return
// 		case ev := <-c.c.Events():
// 			switch e := ev.(type) {
// 			case *kafka.Message:
// 				fmt.Printf("confluent\nkey : %s, Value : %s\n", string(e.Key), string(e.Value))

// 				ad := adapter.KafkaData{
// 					Key: string(e.Key),
// 				}
// 				if err := json.Unmarshal(e.Value, &ad.Value); err != nil {
// 					continue
// 				}
// 				d := adapter.AppToKafka(&ad)
// 				// TODO : check valid
// 				out <- d
// 			case kafka.AssignedPartitions:
// 				fmt.Fprintf(os.Stderr, "%% %v\n", e)
// 				c.c.Assign(e.Partitions)
// 			case kafka.RevokedPartitions:
// 				fmt.Fprintf(os.Stderr, "%% %v\n", e)
// 				c.c.Unassign()
// 			case kafka.PartitionEOF:
// 				fmt.Printf("%% Reached %v\n", e)
// 			case kafka.Error:
// 				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
// 			}
// 		}
// 	}
// }

// func (c *consumer) stop() {
// 	c.ctx <- struct{}{}
// }

// // sensor_1: {"nid" : "node_1", "values" : [0.013, 0.0032], "timestamp" : "2020-08-03 16:20:55"}
