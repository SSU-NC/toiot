package logicCoreUC

import (
	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
	"github.com/KumKeeHyun/PDK/logic-core/domain/repository"
	"github.com/KumKeeHyun/PDK/logic-core/domain/service"
)

type logicCoreUsecase struct {
	mr repository.MetaRepo
	ks service.KafkaConsumerGroup
	es service.ElasticClient
	ls service.LogicCore
}

func NewLogicCoreUsecase(mr repository.MetaRepo, ks service.KafkaConsumerGroup, es service.ElasticClient, ls service.LogicCore) *logicCoreUsecase {
	lcu := &logicCoreUsecase{
		mr: mr,
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

			lchs := lcu.ls.GetLogicChans(rawData.Key)
			if lchs != nil {
				for _, ch := range lchs {
					ch <- ld
				}
			}

			out <- lcu.ToDocument(&ld)
		}
	}()

	return lcu
}

func (lu *logicCoreUsecase) SetLogicChain(r *model.ChainRequest) error {
	// TODO : check chain request validate
	lu.ls.CreateAndStartLogic(r)
	return nil
}

func (lu *logicCoreUsecase) RemoveLogicChain(lname string) error {
	return lu.ls.RemoveLogic(lname)
}
