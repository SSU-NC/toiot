package kcTest

import (
	"fmt"
	"time"

	"github.com/seheee/PDK/logic-core/domain/model"
)

type consumer struct {
	// c *kafka.Consumer
	ctx chan struct{}
}

func NewConsumer() *consumer {
	return &consumer{
		ctx: make(chan struct{}, 1),
	}
}

func (c *consumer) run(out chan<- model.KafkaData) {
	for {
		select {
		case <-c.ctx:
			fmt.Printf("consumer stop\n")
			return
		case <-time.After(3 * time.Second):
			out <- model.KafkaData{
				Key: "sensor_1",
				Value: model.SensorData{
					NID:       "node_1",
					Values:    []float64{0.032, 0.24},
					Timestamp: time.Now(),
				},
			}
		}
	}
}

func (c *consumer) stop() {
	c.ctx <- struct{}{}
}
