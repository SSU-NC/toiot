package elasticContainer

import (
	"bytes"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch"
)

func NewElasticCli() (es *elasticsearch.Client, err error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://220.70.2.1:9200/"},
	}

	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		log.Println(es.Info())
		// => dial tcp: i/o timeout
	}
	return
}

func RunElasticCli(es *elasticsearch.Client, input chan string, e chan int) {
	go func() {
		run := true
		for run == true {
			select {
			case doc := <-input:
				fmt.Printf("push to elk : \n%s\n", doc)
				es.Index(
					"test",
					bytes.NewReader(bytes.NewBufferString(doc).Bytes()),
				)
			case <-e:
				run = false
			}
		}
	}()
}
