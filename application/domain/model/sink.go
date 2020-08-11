package model

import (
	"errors"
	"strconv"
	"strings"
)

type Sink struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"type:varchar(32);unique"`
	Location string `json:"location" gorm:"type:varchar(64)"`
	IP       string `json:"ip" gorm:"type:char(16);not null"`
	Nodes    []Node `json:"nodes" gorm:"foreignkey:sink_id"`
}

func (s *Sink) CheckIP() error {
	sl := strings.Split(s.IP, ":")
	if len(sl) != 2 {
		return errors.New("ip format error:(ip:port)")
	}

	_ip := sl[0]
	ip := strings.Split(_ip, ".")
	if len(ip) != 4 {
		return errors.New("ip format error:(0.0.0.0)")
	}
	for _, n := range ip {
		if num, err := strconv.Atoi(n); err != nil {
			return err
		} else if num < 0 || num > 255 {
			return errors.New("ip format error:(ip range)")

		}
	}
	if _, err := strconv.Atoi(sl[1]); err != nil {
		return err
	}
	return nil
}

func (Sink) TableName() string {
	return "sinks"
}
