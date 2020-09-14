package logicCoreUC

import (
	"strings"

	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
)

func (lcuc *logicCoreUsecase) ToLogicData(kd *model.KafkaData) (model.LogicData, error) {
	n, err := lcuc.rr.FindNode(kd.NodeID)
	if err != nil {
		return model.LogicData{}, err
	}
	s, err := lcuc.rr.FindSensor(kd.SensorID)
	if err != nil {
		return model.LogicData{}, err
	}

	vl := map[string]float64{}
	for i, v := range s.SensorValues {
		vl[v] = kd.Values[i]
	}
	return model.LogicData{
		SensorID:   kd.SensorID,
		SensorName: s.Name,
		Values:     vl,
		Node:       *n,
		Timestamp:  kd.Timestamp,
	}, nil
}

func (lcuc *logicCoreUsecase) ToDocument(ld *model.LogicData) model.Document {
	return model.Document{
		Index: "toiot-" + strings.ReplaceAll(ld.SensorName, " ", "-") + "-" + ld.Node.SinkName,
		Doc:   *ld,
	}
}
