package elasticClient

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
	"github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/elastic/go-elasticsearch"
)

var elasticClient *client

type client struct {
	es *elasticsearch.Client
	in chan model.Document
}

func NewElasticClient() *client {
	if elasticClient != nil {
		return elasticClient
	}

	inBufSize := 100

	config := elasticsearch.Config{
		Addresses: setting.ElasticSetting.Addresses,
	}
	cli, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil
	}

	elasticClient = &client{
		es: cli,
		in: make(chan model.Document, inBufSize),
	}

	go elasticClient.run()

	return elasticClient
}

func (ec *client) run() {
	for doc := range elasticClient.in {
		fmt.Printf("Doc: %v\n", doc)
		d, err := json.Marshal(doc.Doc)
		if err != nil {
			continue
		}
		ec.es.Index(
			doc.Index,
			bytes.NewReader(d),
		)
	}
}

func (ec *client) GetInput() chan<- model.Document {
	if ec != nil {
		return ec.in
	}
	return nil
}
