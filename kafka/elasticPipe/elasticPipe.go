package elasticPipe

import (
	"bytes"
	"encoding/json"

	"github.com/elastic/go-elasticsearch"
)

type setting struct {
	addresses []string
}

var elasticSetting = setting{
	addresses: []string{
		"http://220.70.2.1:9200/",
	},
}

var es *elasticsearch.Client

const BUFSIZE = 1

func Setup() (*elasticsearch.Client, error) {
	config := elasticsearch.Config{
		Addresses: elasticSetting.addresses,
	}
	cli, err := elasticsearch.NewClient(config)
	es = cli
	return es, err
}

func PushToElastic(in <-chan ElasticData) <-chan string {
	out := make(chan string, BUFSIZE)
	go func() {
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
