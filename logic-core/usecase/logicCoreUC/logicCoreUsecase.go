package logicCoreUC

import (
	"github.com/KumKeeHyun/toiot/logic-core/domain/repository"
	"github.com/KumKeeHyun/toiot/logic-core/domain/service"
)

type logicCoreUsecase struct {
	rr repository.RegistRepo
	ks service.KafkaConsumerGroup
	es service.ElasticClient
	ls service.LogicService
}

func NewLogicCoreUsecase(rr repository.RegistRepo,
	ks service.KafkaConsumerGroup,
	es service.ElasticClient,
	ls service.LogicService) *logicCoreUsecase {
	lcu := &logicCoreUsecase{
		rr: rr,
		ks: ks,
		es: es,
		ls: ls,
	}

	in := lcu.ks.GetOutput()
	out := lcu.es.GetInput()

	go func() {
		for rawData := range in {
			ld, err := lcu.ToLogicData(&rawData)
			if err != nil {
				continue
			}

			lchs, err := lcu.ls.GetLogicChans(ld.SensorID)
			if err == nil {
				for _, ch := range lchs {
					ch <- ld
				}
			}
			out <- lcu.ToDocument(&ld)
		}
	}()

	return lcu
}
