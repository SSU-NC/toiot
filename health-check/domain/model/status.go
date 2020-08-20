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

// temp code for alarm
var StateChan chan struct{}

func init() {
	StateChan = make(chan struct{}, 1)
}

type Status struct {
	State       int       `json:"state"`
	Work        bool      `json:"work"`
	Count       int       `json:"count"`
	LastConnect time.Time `json:"last_connect"`
}

func (s *Status) SetState(v int) {
	s.State = v
	// TODO : alarm to react app
	// temp code
	StateChan <- struct{}{}
}

func (s *Status) CheckDrop() bool {
	now := time.Now()
	timeout := s.LastConnect.Add(time.Duration(setting.StatusSetting.Drop) * time.Hour)
	return now.After(timeout)
}

func (s *Status) CheckCnt() {
	if s.Count > 0 {
		s.Count--
	}

	if s.Count == 0 {
		// if it keep in work state for a certain time
		// set state : YELLOW -> GREEN
		// the opposite case set state : YELLOW -> RED
		if s.Work {
			s.SetState(GREEN)
		} else {
			s.SetState(RED)
		}
		//TODO event
	}
}

func (s *Status) Event(work bool, t time.Time) {
	// if work state change intermittently
	// set state : GREEN/RED -> YELLOW
	if s.Work != work {
		s.SetState(YELLOW)
		s.Work = work
		s.Count = setting.StatusSetting.Count
	}
	// Update time for drop check
	if work {
		s.LastConnect = t
	}
}
