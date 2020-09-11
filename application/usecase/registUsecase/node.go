package registUsecase

import "github.com/KumKeeHyun/toiot/application/domain/model"

func (ru *registUsecase) GetSinks() ([]model.Sink, error) {
	return ru.sir.FindsWithTopic()
}

func (ru *registUsecase) RegistSink(s *model.Sink) error {
	return ru.sir.Create(s)
}

func (ru *registUsecase) UnregistSink(s *model.Sink) error {
	return ru.sir.Delete(s)
}

func (ru *registUsecase) GetNodes() ([]model.Node, error) {
	return ru.ndr.FindsWithSensorsValues()
}

func (ru *registUsecase) RegistNode(n *model.Node) error {
	return ru.ndr.Create(n)
}

func (ru *registUsecase) UnregistNode(n *model.Node) error {
	return ru.ndr.Delete(n)
}
