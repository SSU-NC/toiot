package eventUsecase

import (
	"sync"

	"github.com/KumKeeHyun/toiot/application/adapter"
	"github.com/KumKeeHyun/toiot/application/domain/model"
)

func waitRespGroup(path string, body interface{}, ll []model.LogicService) {
	var wg sync.WaitGroup
	for _, l := range ll {
		wg.Add(1)
		go func(_l model.LogicService) {
			url := makeUrl(_l.Addr, path)
			eventClient.R().SetBody(body).Post(url)
		}(l)
	}
	wg.Wait()
}

func (eu *eventUsecase) DeleteSinkEvent(s *model.Sink) error {
	path := "/event/sink/delete"

	ll, err := eu.lsr.FindsByTopicID(s.Topic.ID)
	if err != nil {
		return err
	}

	waitRespGroup(path, s.Nodes, ll)
	// var wg sync.WaitGroup
	// for _, l := range ll {
	// 	wg.Add(1)
	// 	go func() {
	// 		url := makeUrl(l.Addr, path)
	// 		eventClient.R().SetBody(s.Nodes).Post(url)
	// 	}()
	// }
	// wg.Wait()

	return nil
}

func (eu *eventUsecase) CreateNodeEvent(n *model.Node) error {
	path := "/event/node/create"

	ll, err := eu.lsr.FindsByTopicID(n.Sink.Topic.ID)
	if err != nil {
		return err
	}
	waitRespGroup(path, *n, ll)

	return nil
}

func (eu *eventUsecase) DeleteNodeEvent(n *model.Node) error {
	path := "/event/node/delete"

	ll, err := eu.lsr.FindsByTopicID(n.Sink.Topic.ID)
	if err != nil {
		return err
	}
	waitRespGroup(path, *n, ll)

	return nil
}

func (eu *eventUsecase) DeleteSensorEvent(s *model.Sensor) error {
	path := "/event/sensor/delete"

	ll, err := eu.lsr.Finds()
	if err != nil {
		return err
	}
	waitRespGroup(path, *s, ll)

	return nil
}

func (eu *eventUsecase) CreateLogicEvent(l *adapter.Logic) error {
	path := "/event/logic/create"

	ll, err := eu.lsr.Finds()
	if err != nil {
		return err
	}
	waitRespGroup(path, *l, ll)

	return nil
}

func (eu *eventUsecase) DeleteLogicEvent(l *adapter.Logic) error {
	path := "/event/logic/delete"

	ll, err := eu.lsr.Finds()
	if err != nil {
		return err
	}
	waitRespGroup(path, *l, ll)

	return nil
}
