package registUsecase

import "github.com/KumKeeHyun/toiot/application/domain/repository"

type registUsecase struct {
	sir repository.SinkRepo
	ndr repository.NodeRepo
	snr repository.SensorRepo
	lgr repository.LogicRepo
	lsr repository.LogicServiceRepo
	tpr repository.TopicRepo
}

func NewRegistUsecase(sir repository.SinkRepo,
	ndr repository.NodeRepo,
	snr repository.SensorRepo,
	lgr repository.LogicRepo,
	lsr repository.LogicServiceRepo,
	tpr repository.TopicRepo) *registUsecase {
	return &registUsecase{
		sir: sir,
		ndr: ndr,
		snr: snr,
		lgr: lgr,
		lsr: lsr,
		tpr: tpr,
	}
}
