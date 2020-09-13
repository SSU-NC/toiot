package registUsecase

import (
	"github.com/KumKeeHyun/toiot/application/adapter"
	"github.com/KumKeeHyun/toiot/application/domain/model"
)

func (ru *registUsecase) GetSinks() ([]model.Sink, error) {
	return ru.sir.FindsWithTopic()
}

func (ru *registUsecase) GetSinksByTopicID(tid int) ([]model.Sink, error) {
	return ru.sir.FindsByTopicIDWithNodesSensorsValuesLogics(tid)
}

func (ru *registUsecase) RegistSink(s *model.Sink) error {
	return ru.sir.Create(s)
}

func (ru *registUsecase) UnregistSink(s *model.Sink) error {
	return ru.sir.Delete(s)
}

func (ru *registUsecase) GetPageCount(size int) int {
	return ru.ndr.GetPages(size)
}

func (ru *registUsecase) GetNodes() ([]model.Node, error) {
	return ru.ndr.FindsWithSensorsValues()
}

func (ru *registUsecase) GetNodesPage(p adapter.Page) ([]model.Node, error) {
	return ru.ndr.FindsPage(p)
}

func (ru *registUsecase) GetNodesSquare(sq adapter.Square) ([]model.Node, error) {
	return ru.ndr.FindsSquare(sq)
}

func (ru *registUsecase) RegistNode(n *model.Node) error {
	return ru.ndr.Create(n)
}

func (ru *registUsecase) UnregistNode(n *model.Node) error {
	return ru.ndr.Delete(n)
}
