package model

import (
	"time"
)

const (
	RED    = 0 // 미동작
	YELLOW = 1 // 형태 바뀔 경우 가운데서 중재 단계
	GREEN  = 2 // 동작
)

type SinkStatus struct {
	SinkID  int          `json:"sid"`
	Satates []NodeStatus `json:"states"`
}

type NodeStatus struct {
	NodeID  int `json:"nid"`
	State   int `json:"state"`
	Battery int `json:"battery"`
}

type Status struct {
	State       int       `json:"state"`
	Work        bool      `json:"work"`
	LastConnect time.Time `json:"last_connect"`
}

func NewStatus(work bool, t time.Time) Status { // 인자로 받은 work 여부로 Status 구조체 설정
	res := Status{
		Work:        work,
		LastConnect: t,
	}
	if work {
		res.State = GREEN
	} else {
		res.State = RED
	}
	return res
}

func (s *Status) setState(v int) { // 인자로 받은 v로 Status구조체 변경
	s.State = v
	switch v {
	case RED:
		s.Work = false
	case GREEN:
		s.Work = true
	case YELLOW:
		s.Work = !s.Work
	}
}

func (s *Status) UpdateState(work bool, t time.Time) bool {
	isChange := false
	// Update time for drop
	if work {
		s.LastConnect = t
	}
	if s.State == YELLOW {
		if work {
			s.setState(GREEN)
		} else {
			s.setState(RED)
		}
		isChange = true
	} else if s.Work != work {
		s.setState(YELLOW)
		isChange = true
	}
	return isChange
}
func (s *Status) CheckDrop() bool {
	s.setState(RED)
	now := time.Now()
	timeout := time.Now() //s.LastConnect.Add(time.Duration(setting.StatusSetting.Drop) * time.Hour)
	return now.After(timeout)
}
