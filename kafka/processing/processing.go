package processing

import (
	"github.com/KumKeeHyun/PDK/kafka/elasticPipe"
	"github.com/KumKeeHyun/PDK/kafka/kafkaPipe"
)

const BUFSIZE = 1

var valueNames = []string{
	"x",
	"y",
	"z",
}

func ProcessingPipe(in <-chan kafkaPipe.KafkaData) <-chan elasticPipe.ElasticData {
	out := make(chan elasticPipe.ElasticData, BUFSIZE)
	go func() {
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
	values := in.Value["value"].([]interface{})
	newValues := map[string]interface{}{}
	for i, vn := range valueNames {
		newValues[vn] = values[i]
	}
	out.Doc["value"] = newValues

	return out, nil
}
