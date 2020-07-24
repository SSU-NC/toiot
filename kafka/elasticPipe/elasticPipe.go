package elasticPipe

import (
	"bytes"
	"encoding/json"

	"github.com/KumKeeHyun/PDK/kafka/setting"
	"github.com/elastic/go-elasticsearch"
)

var es *elasticsearch.Client

const BUFSIZE = 100

func Setup() (*elasticsearch.Client, error) {
	config := elasticsearch.Config{
		Addresses: setting.ElasticSetting.Addresses,
	}
	cli, err := elasticsearch.NewClient(config)
	es = cli
	return es, err
}

func PushToElastic(in <-chan ElasticData) <-chan string {
	out := make(chan string, BUFSIZE)
	go func() {
		defer func() {
			close(out)
		}()
		for data := range in {
			doc, err := json.Marshal(data.Doc)
			if err != nil {
				continue
			}
			es.Index(
				data.Index,
				bytes.NewReader(doc),
			)
			out <- string(doc)
		}
	}()
	return out
}
