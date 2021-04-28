package eventUsecase

import (
	"log"
	"sync"

	"github.com/KumKeeHyun/toiot/application/adapter"
	"github.com/KumKeeHyun/toiot/application/domain/model"
)

func waitRespGroup(e EVENT, body interface{}, ll []model.LogicService) (prl []pingRequest) {
	var wg sync.WaitGroup
	for _, l := range ll {
		wg.Add(1)
		go func(_l model.LogicService) {
			url := makeUrl(_l.Addr, EventPath[e])
			resp, _ := eventClient.R().SetBody(body).Post(url)
			log.Println("Post 내용 : ", body, "url : ", url)
			if !resp.IsSuccess() {
				prl = append(prl, pingRequest{_l, e, body})
			}
			wg.Done()
		}(l)
	}
	wg.Wait()
	return
}

func (eu *eventUsecase) DeleteSinkEvent(s *model.Sink) error {
	e := DeleteSink

	ll, err := eu.lsr.FindsByTopicID(s.Topic.ID)
	if err != nil {
		return err
	}

	eu.requestRetry = append(eu.requestRetry, waitRespGroup(e, s.Nodes, ll)...)
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
func (eu *eventUsecase) CreateSinkEvent(s *model.Sink) error {
	e := CreateSink
	sinkaddr := adapter.SinkAddr{
		Sid:  s.ID,
		Addr: s.Addr,
	}

	ll, err := eu.lsr.FindsByTopicID(s.Topic.ID)
	if err != nil {
		return err
	}
	eu.requestRetry = append(eu.requestRetry, waitRespGroup(e, sinkaddr, ll)...)

	return nil
}

func (eu *eventUsecase) CreateNodeEvent(n *model.Node) error {
	e := CreateNode

	ll, err := eu.lsr.FindsByTopicID(n.Sink.Topic.ID)
	if err != nil {
		return err
	}
	eu.requestRetry = append(eu.requestRetry, waitRespGroup(e, *n, ll)...)

	return nil
}

func (eu *eventUsecase) DeleteNodeEvent(n *model.Node) error {
	e := DeleteNode

	ll, err := eu.lsr.FindsByTopicID(n.Sink.Topic.ID)
	if err != nil {
		return err
	}
	eu.requestRetry = append(eu.requestRetry, waitRespGroup(e, *n, ll)...)

	return nil
}

func (eu *eventUsecase) DeleteSensorEvent(s *model.Sensor) error {
	e := DeleteSensor

	ll, err := eu.lsr.Finds()
	if err != nil {
		return err
	}
	eu.requestRetry = append(eu.requestRetry, waitRespGroup(e, *s, ll)...)

	return nil
}

func (eu *eventUsecase) CreateLogicEvent(l *model.Logic) error {
	e := CreateLogic

	ll, err := eu.lsr.Finds()
	if err != nil {
		return err
	}
	eu.requestRetry = append(eu.requestRetry, waitRespGroup(e, *l, ll)...)

	return nil
}

func (eu *eventUsecase) DeleteLogicEvent(l *model.Logic) error {
	e := DeleteLogic

	ll, err := eu.lsr.Finds()
	if err != nil {
		return err
	}
	eu.requestRetry = append(eu.requestRetry, waitRespGroup(e, *l, ll)...)

	return nil
}
