package adapter

import (
	"encoding/json"

	"github.com/KumKeeHyun/toiot/application/domain/model"
)

type Logic struct {
	ID        int          `json:"id"`
	LogicName string       `json:"logic_name"`
	Elems     []Element    `json:"elems"` // 엘렘 배열 , value time email 중 선택한거 드가있음
	SensorID  int          `json:"sensor_id"`
	Sensor    model.Sensor `json:"sensor"`
}

/*
{
  "aid": int,
  "value": int,
  "sleep": int
}
*/
type Element struct {
	Elem string                 `json:"elem"` //value time email 중 선택한거
	Arg  map[string]interface{} `json:"arg"`  //거기에 맞는 A
}

func LogicToAdapter(ml *model.Logic) (Logic, error) {
	var elems []Element
	if err := json.Unmarshal([]byte(ml.Elems), &elems); err != nil {
		return Logic{}, err
	} else {
		return Logic{
			ID:        ml.ID,
			LogicName: ml.Name,
			Elems:     elems,
			SensorID:  ml.SensorID,
			Sensor:    ml.Sensor,
		}, nil
	}
}

func LogicsToAdapter(mll []model.Logic) []Logic {
	res := make([]Logic, 0, len(mll))
	for _, ml := range mll {
		if l, err := LogicToAdapter(&ml); err == nil {
			res = append(res, l)
		}
	}
	return res
}

func LogicToModel(l *Logic) (model.Logic, error) {
	if b, err := json.Marshal(l.Elems); err != nil {
		return model.Logic{}, err
	} else {
		return model.Logic{
			ID:       l.ID,
			Name:     l.LogicName,
			Elems:    string(b),
			SensorID: l.SensorID,
		}, nil
	}
}

func LogicsToModel(ll []Logic) []model.Logic {
	res := make([]model.Logic, 0, len(ll))
	for _, l := range ll {
		if ml, err := LogicToModel(&l); err == nil {
			res = append(res, ml)
		}
	}
	return res
}
