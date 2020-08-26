package sinkUsecase

import (
	"github.com/seheee/PDK/application/domain/model"
	"github.com/seheee/PDK/application/domain/repository"
)

type sinkUsecase struct {
	sir repository.SinkRepository
	nr  repository.NodeRepository
}

func NewSinkUsecase(sir repository.SinkRepository, nr repository.NodeRepository) *sinkUsecase {
	return &sinkUsecase{
		sir: sir,
		nr:  nr,
	}
}

func (siu *sinkUsecase) GetAllSinks() ([]model.Sink, error) {
	sis, err := siu.sir.GetAll()
	if err != nil {
		return nil, err
	}
	return sis, nil
}

func (siu *sinkUsecase) GetAllSinksWithNodes() ([]model.Sink, error) {
	sis, err := siu.sir.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range sis {
		if sis[i].Nodes, err = siu.nr.GetBySinkID(sis[i].ID); err != nil {
			return nil, err
		}
	}
	return sis, nil
}

func (siu *sinkUsecase) GetSinkByID(sinkID uint) (*model.Sink, error) {
	si, err := siu.sir.GetByID(sinkID)
	if err != nil {
		return nil, err
	}
	return si, nil
}

func (siu *sinkUsecase) GetSinkByIDWithNodes(sinkID uint) (*model.Sink, error) {
	si, err := siu.sir.GetByID(sinkID)
	if err != nil {
		return nil, err
	}
	si.Nodes, err = siu.nr.GetBySinkID(sinkID)
	if err != nil {
		return nil, err
	}
	return si, nil
}

func (siu *sinkUsecase) RegisterSink(si *model.Sink) (*model.Sink, error) {
	err := siu.sir.Create(si)
	if err != nil {
		return nil, err
	}
	return si, nil
}

func (siu *sinkUsecase) DeleteSink(si *model.Sink) error {
	return siu.sir.Delete(si)
}
