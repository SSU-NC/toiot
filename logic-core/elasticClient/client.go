package elasticClient

import (
	"fmt"
	"strings"
	"time"

	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
	"github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/elastic/go-elasticsearch/v8"
)

var elasticClient *client

type client struct {
	es *elasticsearch.Client
	in chan model.Document

	ticker  *time.Ticker
	docBuf  []*model.Document
	bufSize int
}

func NewElasticClient() *client {
	if elasticClient != nil {
		return elasticClient
	}
	inBufSize := 100

	config := elasticsearch.Config{
		Addresses:  setting.ElasticSetting.Addresses,
		MaxRetries: 3,
	}
	cli, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil
	}

	elasticClient = &client{
		es:      cli,
		in:      make(chan model.Document, inBufSize),
		ticker:  time.NewTicker(time.Duration(2) * time.Second),
		docBuf:  make([]*model.Document, 0, 100),
		bufSize: 90,
	}

	go elasticClient.run()

	return elasticClient
}

func (ec *client) run() {
	for {
		select {
		case doc := <-ec.in:
			ec.insertDoc(&doc)
		case <-ec.ticker.C:
			ec.bulk()
		}
	}
}

func (ec *client) GetInput() chan<- model.Document {
	if ec != nil {
		return ec.in
	}
	return nil
}

func (ec *client) insertDoc(d *model.Document) {
	ec.docBuf = append(ec.docBuf, d)
	if len(ec.docBuf) >= ec.bufSize {
		ec.bulk()
	}
}

func (ec *client) bulk() {
	if len(ec.docBuf) > 0 {
		bulkStr := strings.Join(docsToSlice(ec.docBuf), "")
		res, _ := ec.es.Bulk(strings.NewReader(bulkStr))
		res.Body.Close()

		fmt.Println(bulkStr)
		ec.docBuf = make([]*model.Document, 0, 100)
	}
}

func docsToSlice(docs []*model.Document) []string {
	res := make([]string, len(docs))
	for i, doc := range docs {
		res[i] = doc.String()
	}
	return res
}
