package service

import "github.com/KumKeeHyun/PDK/logic-core/domain/model"

type ElasticClient interface {
	GetInput() chan<- model.Document
}
