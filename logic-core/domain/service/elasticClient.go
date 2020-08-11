package service

import "github.com/seheee/PDK/logic-core/domain/model"

type ElasticClient interface {
	GetInput() chan<- model.Document
}
