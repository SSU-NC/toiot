package elasticClient

import (
	"fmt"
	"strings"
	"time"

	"github.com/KumKeeHyun/PDK/health-check/adapter.go"
	"github.com/KumKeeHyun/PDK/health-check/setting"
	"github.com/elastic/go-elasticsearch/v8"
)

var elasticClient *client

type client struct {
	es *elasticsearch.Client
	in chan adapter.Document

	ticker  *time.Ticker
	docBuf  []*adapter.Document
	bufSize int
}

func NewElasticClient() *client {
	if elasticClient != nil {
		return elasticClient
	}
	inBufSize := 100

	config := elasticsearch.Config{
		Addresses: setting.ElasticSetting.Addresses,
		//MaxRetries: 3,
	}
	cli, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil
	}

	elasticClient = &client{
		es:      cli,
		in:      make(chan adapter.Document, inBufSize),
		ticker:  time.NewTicker(10 * time.Second),
		docBuf:  make([]*adapter.Document, 0, 100),
		bufSize: 100,
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

func (ec *client) GetInput() chan<- adapter.Document {
	if ec != nil {
		return ec.in
	}
	return nil
}

func (ec *client) insertDoc(d *adapter.Document) {
	ec.docBuf = append(ec.docBuf, d)
	if len(ec.docBuf) >= (ec.bufSize - 10) {
		ec.bulk()
	}
}

func (ec *client) bulk() {
	if len(ec.docBuf) > 0 {
		bulkStr := strings.Join(docsToSlice(ec.docBuf), "")
		fmt.Printf("bulk\n%s", bulkStr)
		res, err := ec.es.Bulk(strings.NewReader(bulkStr))
		if res != nil && err == nil {
			res.Body.Close()
		} else {
			fmt.Printf("bulk error : %s\n", err.Error())
		}

		ec.docBuf = make([]*adapter.Document, 0, ec.bufSize)
	}
}

func docsToSlice(docs []*adapter.Document) []string {
	res := make([]string, 0, len(docs))
	for _, doc := range docs {
		res = append(res, doc.String())
	}
	return res
}
