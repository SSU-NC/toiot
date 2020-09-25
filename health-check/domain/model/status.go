package model

import (
	"time"

	"github.com/KumKeeHyun/PDK/health-check/setting"
)

const (
	RED    = 0
	YELLOW = 1
	GREEN  = 2
)

type NodeStatus struct {
	NodeID int `json:"nid"`
	State  int `json:"state"`
}

type Status struct {
	State       int       `json:"state"`
	Work        bool      `json:"work"`
	Count       int       `json:"count"`
	LastConnect time.Time `json:"last_connect"`
}

func NewStatus(work bool, t time.Time) Status {
	res := Status{
		Work:        work,
		Count:       -1,
		LastConnect: t,
	}
	if work {
		res.State = GREEN
	} else {
		res.State = RED
	}
	return res
}

func (s *Status) setState(v int) {
	s.State = v
	switch v {
	case RED:
		s.Count = -1
		s.Work = false
	case GREEN:
		s.Count = -1
		s.Work = true
	case YELLOW:
		s.Count = setting.StatusSetting.Count
		s.Work = !s.Work
	}
}

func (s *Status) decreaseCnt() {
	if s.Count >= 0 {
		s.Count--
	}
}

func (s *Status) UpdateState(work bool, t time.Time) bool {
	isChange := false
	// Update time for drop
	if work {
		s.LastConnect = t
	}
	if s.Work != work {
		s.setState(YELLOW)
		isChange = true
	} else {
		s.decreaseCnt()
		if s.Count == 0 {
			if s.Work {
				s.setState(GREEN)
			} else {
				s.setState(RED)
			}
			isChange = true
		}
	}
	return isChange
}

func (s *Status) CheckDrop() bool {
	s.setState(RED)
	now := time.Now()
	timeout := s.LastConnect.Add(time.Duration(setting.StatusSetting.Drop) * time.Hour)
	return now.After(timeout)
}
