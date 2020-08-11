package logicCoreUC

import (
	"errors"
	"github.com/seheee/PDK/logic-core/domain/model"
	"github.com/seheee/PDK/logic-core/domain/repository"
	"github.com/seheee/PDK/logic-core/domain/service"
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

func (lu *logicCoreUsecase) SetLogicChain(r *model.RingRequest) error {
	// TODO : check chain request validate
	_, err := lu.mr.GetSensor(r.Sensor)
	if err != nil {
		return errors.New("sensor does not exist")
	}
	chs := lu.ls.GetLogicChans(r.Sensor)
	_, ok := chs[r.LogicName]
	if ok {
		return errors.New("logic name already exists")
	}
	go lu.ls.CreateAndStartLogic(r)
	return nil
}

func (lu *logicCoreUsecase) RemoveLogicChain(lname string) error {
	return lu.ls.RemoveLogic(lname)
}

func (lu *logicCoreUsecase) RemoveLogicChainsBySID(sid string) error {
	return lu.ls.RemoveLogicsBySID(sid)
}


