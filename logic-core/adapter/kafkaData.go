package adapter

import (
	"time"

	"github.com/seheee/PDK/logic-core/domain/model"
)

type SensorData struct {
	NID       string    `json:"nid"`
	Values    []float64 `json:"values"`
	Timestamp string    `json:"timestamp"`
}

type KafkaData struct {
	Key   string     `json:"key"`
	Value SensorData `json:"value"`
}

func AppToKafka(kd *KafkaData) model.KafkaData {
	t, err := time.Parse("2006-01-02 15:04:05", kd.Value.Timestamp)
	if err != nil {
		return model.KafkaData{}
	}
	return model.KafkaData{
		Key: kd.Key,
		Value: model.SensorData{
			NID:       kd.Value.NID,
			Values:    kd.Value.Values,
			Timestamp: t,
		},
	}
}
