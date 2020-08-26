package esTest

import (
	"fmt"

	"github.com/seheee/PDK/logic-core/domain/model"
)

var elasticClient *client

type client struct {
	// es *elastic.Client
	in chan model.Document
}

func NewElasticClient() *client {
	if elasticClient != nil {
		return elasticClient
	}

	inBufSize := 100

	elasticClient = &client{
		in: make(chan model.Document, inBufSize),
	}
	go func() {
		for doc := range elasticClient.in {
			fmt.Printf("Doc: %v\n", doc)
		}
	}()
	return elasticClient
}

func (ec *client) GetInput() chan<- model.Document {
	if ec != nil {
		return ec.in
	}
	return nil
}
