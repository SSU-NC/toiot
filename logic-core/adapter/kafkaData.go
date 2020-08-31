package adapter

import (
	"time"

	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
)

var loc *time.Location

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
	t, err := time.ParseInLocation("2006-01-02 15:04:05", kd.Value.Timestamp, loc)
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

func init() {
	loc, _ = time.LoadLocation("Asia/Seoul")
}
