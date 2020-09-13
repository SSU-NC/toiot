package service

import "github.com/KumKeeHyun/toiot/logic-core/domain/model"

type ElasticClient interface {
	GetInput() chan<- model.Document
}
