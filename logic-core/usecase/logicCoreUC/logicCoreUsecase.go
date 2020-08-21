package logicCoreUC

import (
	"errors"

	"github.com/seheee/PDK/logic-core/domain/model"
	"github.com/seheee/PDK/logic-core/domain/repository"
	"github.com/seheee/PDK/logic-core/domain/service"
	
)

type logicCoreUsecase struct {
	mr repository.MetaRepo
	lr repository.LogicRepo
	ks service.KafkaConsumerGroup
	es service.ElasticClient
	ls service.LogicCore
	event chan interface{}
}

func NewLogicCoreUsecase(mr repository.MetaRepo, lr repository.LogicRepo, ks service.KafkaConsumerGroup, es service.ElasticClient, ls service.LogicCore, event chan interface{}) *logicCoreUsecase {
	lcu := &logicCoreUsecase{
		mr: mr,
		lr: lr,
		ks: ks,
		es: es,
		ls: ls,
		event: event,
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
	/*_, err := lu.mr.GetSensor(r.Sensor)
	if err != nil {
		return errors.New("sensor does not exist")
	}*/
	chs := lu.ls.GetLogicChans(r.Sensor)
	_, ok := chs[r.LogicName]
	if ok {
		return errors.New("logic name already exists")
	}
	id, err := lu.lr.Create(r)
	if err != nil {
		return err
	}
	go lu.ls.CreateAndStartLogic(r, id, lu.event)
	return nil
}

func (lu *logicCoreUsecase) RemoveLogicChain(id string) error {
	if err := lu.lr.Delete(id); err != nil {
		return err
	}
	return lu.ls.RemoveLogic(id)
}

func (lu *logicCoreUsecase) RemoveLogicChainsBySID(sid string) error {
	return lu.ls.RemoveLogicsBySID(sid)
}

func (lu *logicCoreUsecase) GetAllLogics() ([]model.Ring, error) {
	lg, err := lu.lr.GetAll()
	if err != nil {
		return nil, err
	}

	return lg, err
}

