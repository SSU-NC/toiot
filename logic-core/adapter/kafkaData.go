package adapter

import (
	"time"

	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
)

var (
	loc     *time.Location
	timeFmt string
)

func init() {
	loc, _ = time.LoadLocation("Asia/Seoul")
	timeFmt = "2006-01-02 15:04:05"
}

type KafkaData struct {
	SensorID  int       `json:"sensor_id"`
	NodeID    int       `json:"node_id"`
	Values    []float64 `json:"values"`
	Timestamp string    `json:"timestamp"`
}

func KafkaToModel(d *KafkaData) (model.KafkaData, error) {
	t, err := time.ParseInLocation(timeFmt, d.Timestamp, loc)
	if err != nil {
		return model.KafkaData{}, err
	}
	return model.KafkaData{
		SensorID:  d.SensorID,
		NodeID:    d.NodeID,
		Values:    d.Values,
		Timestamp: t,
	}, nil
}
