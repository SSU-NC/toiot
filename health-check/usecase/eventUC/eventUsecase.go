package eventUC

import (
	"strings"
	"time"

	"github.com/KumKeeHyun/PDK/health-check/adapter.go"
	"github.com/KumKeeHyun/PDK/health-check/domain/repository"
	"github.com/KumKeeHyun/PDK/health-check/domain/service"
)

type eventUcsecase struct {
	sr repository.StatusRepo
	ks service.KafkaConsumer
	es service.ElasticClient
}

func NewEventUsecase(sr repository.StatusRepo, ks service.KafkaConsumer, es service.ElasticClient) *eventUcsecase {
	eu := &eventUcsecase{
		sr: sr,
		ks: ks,
		es: es,
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
			for _, n := range ns {
				s, ok := nm[n.UUID]
				if !ok {
					// TODO
					continue
				}
				status, err := eu.sr.Get(n.UUID)
				if err != nil {
					// TODO
					continue
				}
				status.Event(s.State, StrToTime(states.Timestamp))
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

			eu.sr.EndAtomic()
		}
	}()

	return eu
}

func StrToTime(s string) time.Time {
	f := "2006-01-02 15:04:05"
	res, _ := time.Parse(f, s)
	return res
}

func TimeToStr(t time.Time) string {
	return t.String()[:19]
}
