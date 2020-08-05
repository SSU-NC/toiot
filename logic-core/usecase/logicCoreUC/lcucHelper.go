package logicCoreUC

import (
	"strings"

	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
)

func (lcuc *logicCoreUsecase) ToLogicData(kd *model.KafkaData) (model.LogicData, error) {
	n, err := lcuc.mr.GetNode(kd.Value.NID)
	if err != nil {
		return model.LogicData{}, err
	}
	s, err := lcuc.mr.GetSensor(kd.Key)
	if err != nil {
		return model.LogicData{}, err
	}

	v := map[string]float64{}
	for i, vn := range s.ValueNames {
		v[vn] = kd.Value.Values[i]
	}

	return model.LogicData{
		SID:       kd.Key,
		SName:     s.Name,
		Values:    v,
		NodeInfo:  *n,
		Timestamp: kd.Value.Timestamp,
	}, nil
}

func (lcuc *logicCoreUsecase) ToDocument(ld *model.LogicData) model.Document {
	return model.Document{
		Index: strings.ReplaceAll(ld.SName, " ", "-") + "-" + ld.NodeInfo.Group,
		Doc:   *ld,
	}
}
