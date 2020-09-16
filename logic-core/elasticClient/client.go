package elasticClient

import (
	"strings"
	"time"

	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
	"github.com/KumKeeHyun/toiot/logic-core/setting"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-resty/resty/v2"
)

var (
	elasticClient *client
	template      = `
{
	"index_patterns": [
	  "toiot*"
	],
	"settings": {
	  "number_of_shards": 1
	},
	"mappings" : {
	  "properties" : {
		"name" : {
		  "type" : "keyword"
		},
		"node" : {
		  "properties" : {
			"sink_name" : {
			  "type" : "keyword"
			},
			"location" : {
			  "type": "geo_point"
			},
			"name" : {
			  "type" : "keyword"
			}
		  }
		},
		"sensor_id" : {
		  "type" : "long"
		},
		"sensor_name" : {
		  "type" : "keyword"
		},
		"timestamp" : {
		  "type" : "date"
		}
	  }
	}
}
`
)

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
	inBufSize := setting.Elasticsetting.ChanBufSize

	config := elasticsearch.Config{
		Addresses:  setting.Elasticsetting.Addresses,
		MaxRetries: setting.Elasticsetting.RequestRetry,
	}
	cli, err := elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}

	putTemplate := resty.New()

	putTemplate.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(template)).
		Put("http://220.70.2.1:9200/_template/template_1")

	elasticClient = &client{
		es:      cli,
		in:      make(chan model.Document, inBufSize),
		ticker:  time.NewTicker(time.Duration(setting.Elasticsetting.BatchTicker) * time.Second),
		docBuf:  make([]*model.Document, 0, setting.Elasticsetting.BatchSize),
		bufSize: setting.Elasticsetting.BatchSize,
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

/*
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
*/
func (ec *client) GetInput() chan<- model.Document {
	if ec != nil {
		return ec.in
	}
	return nil
}

func (ec *client) insertDoc(d *model.Document) {
	ec.docBuf = append(ec.docBuf, d)
	if len(ec.docBuf) >= (ec.bufSize - 10) {
		ec.bulk()
	}
}

func (ec *client) bulk() {
	if len(ec.docBuf) > 0 {
		bulkStr := strings.Join(docsToSlice(ec.docBuf), "")
		// debug
		//fmt.Printf("elastic : %v\n", bulkStr)

		res, _ := ec.es.Bulk(strings.NewReader(bulkStr))
		res.Body.Close()
		ec.docBuf = make([]*model.Document, 0, ec.bufSize)
	}
}

func docsToSlice(docs []*model.Document) []string {
	res := make([]string, len(docs))
	for i, doc := range docs {
		res[i] = doc.String()
	}
	return res
}
