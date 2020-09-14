package registUsecase

import (
	"github.com/KumKeeHyun/toiot/application/adapter"
	"github.com/KumKeeHyun/toiot/application/domain/model"
)

func (ru *registUsecase) GetSinkPageCount(size int) int {
	return ru.sir.GetPages(size)
}

func (ru *registUsecase) GetSinks() ([]model.Sink, error) {
	return ru.sir.FindsWithTopic()
}

func (ru *registUsecase) GetSinksPage(p adapter.Page) ([]model.Sink, error) {
	return ru.sir.FindsPage(p)
}

func (ru *registUsecase) GetSinksByTopicID(tid int) ([]model.Sink, error) {
	return ru.sir.FindsByTopicIDWithNodesSensorsValuesLogics(tid)
}

func (ru *registUsecase) GetSinkByID(sid int) (*model.Sink, error) {
	return ru.sir.FindByIDWithNodesSensorsValuesTopic(sid)
}

func (ru *registUsecase) RegistSink(s *model.Sink) error {
	return ru.sir.Create(s)
}

func (ru *registUsecase) UnregistSink(s *model.Sink) error {
	return ru.sir.Delete(s)
}

func (ru *registUsecase) GetNodePageCount(p adapter.Page) int {
	return ru.ndr.GetPages(p)
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
