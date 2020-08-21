package eventUC

import (
	"strings"
	"time"

	"github.com/KumKeeHyun/PDK/health-check/domain/model"
	"github.com/KumKeeHyun/PDK/health-check/setting"

	"github.com/KumKeeHyun/PDK/health-check/adapter.go"
	"github.com/KumKeeHyun/PDK/health-check/domain/repository"
	"github.com/KumKeeHyun/PDK/health-check/domain/service"
)

type eventUsecase struct {
	sr    repository.StatusRepo
	ks    service.KafkaConsumer
	es    service.ElasticClient
	event chan struct{}
}

func NewEventUsecase(sr repository.StatusRepo, ks service.KafkaConsumer, es service.ElasticClient, e chan struct{}) *eventUsecase {
	eu := &eventUsecase{
		sr:    sr,
		ks:    ks,
		es:    es,
		event: e,
	}

	in := eu.ks.GetOutput()
	out := eu.es.GetInput()

	go func() {
		for states := range in {
			eu.sr.StartAtomic()

			nm := states.GetNodeMap()
			ns, err := adapter.GetNodeInfo(&states)
			if err != nil {
				// TODO
				continue
			}

			change := false
			for _, n := range ns {
				s, _ := nm[n.UUID]
				status, err := eu.GetNodeStatus(s, states.Timestamp)
				if err != nil {
					// TODO
					continue
				}
				change = status.Event(s.State, StrToTime(states.Timestamp))
				eu.sr.Update(n.UUID, status)

				out <- adapter.Document{
					Index: "hc-" + strings.ReplaceAll(n.Group, " ", "-"),
					Doc: adapter.StateDoc{
						Node:      n,
						Status:    status,
						Timestamp: states.Timestamp,
					},
				}
			}
			if change {
				eu.event <- struct{}{}
			}

			eu.sr.EndAtomic()
		}
	}()

	return eu
}

func (eu *eventUsecase) GetNodeStatus(ns adapter.NodeState, t string) (model.Status, error) {
	res, err := eu.sr.Get(ns.NodeID)
	if err != nil {
		res = model.Status{
			Work:        ns.State,
			Count:       setting.StatusSetting.Count,
			LastConnect: StrToTime(t),
		}
		res.SetState(model.YELLOW)
		eu.event <- struct{}{}

		if err := eu.sr.Create(ns.NodeID, res); err != nil {
			return model.Status{}, err
		}
	}
	return res, nil
}

func StrToTime(s string) time.Time {
	f := "2006-01-02 15:04:05"
	res, _ := time.Parse(f, s)
	return res
}

func TimeToStr(t time.Time) string {
	return t.String()[:19]
}
