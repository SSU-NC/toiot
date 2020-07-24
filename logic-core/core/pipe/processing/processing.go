package processing

import (
	"errors"
	"fmt"

	"github.com/KumKeeHyun/PDK/logic-core/core/pipe/elasticPipe"
	"github.com/KumKeeHyun/PDK/logic-core/core/pipe/kafkaPipe"
	"github.com/KumKeeHyun/PDK/logic-core/logic-core-api/model"
)

const BUFSIZE = 100

func ProcessingPipe(in <-chan kafkaPipe.KafkaData) <-chan elasticPipe.ElasticData {
	out := make(chan elasticPipe.ElasticData, BUFSIZE)
	go func() {
		defer func() {
			close(out)
		}()
		for data := range in {
			res, err := DataProcessing(data)
			if err != nil {
				continue
			}
			out <- res
		}
	}()
	return out
}

func DataProcessing(in kafkaPipe.KafkaData) (elasticPipe.ElasticData, error) {
	out := elasticPipe.ElasticData{
		Index: in.Key,
		Doc:   in.Value,
	}

	v, ok := in.Value["node_uuid"]
	if ok {
		delete(out.Doc, "node_uuid")
		nID := v.(string)
		node, err := model.RegisterRepo.GetNode(nID)
		if err == nil {
			out.Doc["node"] = node
		} else {
			return out, err
		}
	}

	sensor, err := model.RegisterRepo.GetSensor(in.Key)
	if err != nil {
		s := fmt.Sprintf("not exist sensor %s\n", in.Key)
		return out, errors.New(s)
	}

	values := in.Value["value"].([]interface{})
	newValues := map[string]interface{}{}
	for i, vn := range sensor.ValueList {
		newValues[vn] = values[i]
	}
	out.Doc["value"] = newValues

	return out, nil
}
