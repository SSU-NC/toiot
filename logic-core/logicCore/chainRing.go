package logicCore

import (
	"fmt"
	"time"
	"net/smtp"

	"github.com/seheee/PDK/logic-core/domain/model"
)

type rangeRing struct {
	baseRing
	Value string `json:"value"`
	Range []struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"range"`

}
func (r *rangeRing) exec(d *model.LogicData) {
	isRange := func(val float64) bool {
		for _, rg := range r.Range{
			if val < rg.Min && val > rg.Max {
				return false
			}
		}
		return true
	}
	v, ok := d.Values[r.Value]
	if !ok {
		return
	}
	if isRange(v) {
		r.next.exec(d)
	} 
}

type timeRing struct {
	baseRing
	Range []struct {
		Start string `json:"start"`
		End string `json:"end"`
	} `json:"range"`
	
}
func (r *timeRing) exec(d *model.LogicData) {
	isTime := func(ts time.Time) bool {
		for _, rg := range r.Range{
			st, _ := time.Parse("15:04:05", rg.Start)
			et, _ := time.Parse("15:04:05", rg.End)
			if !(ts.After(st) && ts.Before(et)) {
				return false
			}
		}
		return true
	}
	fmt.Println("timeRing")
	ts := d.Timestamp
	ts, _ = time.Parse("15:04:05", ts.Format("15:04:05"))
	if isTime(ts) {
		r.next.exec(d)
	}
}

type groupRing struct {
	baseRing
	Group []string `json:"group"`
}
func (r *groupRing) exec(d *model.LogicData) {
	for _, group := range r.Group {
		if group == d.NodeInfo.Group {
			r.next.exec(d)
		}
	}
}

type emailRing struct {
	baseRing
	Email string `json:"text"`
	Time bool
}
func (r *emailRing) exec(d *model.LogicData) {
	
	if r.Time == true {
		from := "sehee4010@gmail.com"
		pass := ""
		to := r.Email

		body := "sensor \"" + d.SName + "\"" +
				" on node \"" + d.NodeInfo.Name +"\"" 
		
		msg := 	"From: " + from + "\n" +
				"To: " + to + "\n" +
				"Subject: PDK email\n\n" +
				body

		err := smtp.SendMail("smtp.gmail.com:587",
				smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
				from, []string{to}, []byte(msg))

		if err != nil {
			fmt.Printf("smtp error: %s", err)
		}
		fmt.Printf("[email] send to %s when %s\n", r.Email, d.Timestamp)
		t1 := time.NewTimer(time.Second * 180)
		r.Time = false
		go func() {
			<- t1.C
			r.Time = true
			fmt.Println("email timer expired")
		}()
	}
	if r.next != nil {
		r.next.exec(d)
	}
}

type alarmRing struct {
	baseRing
	ch chan interface{}
	Message string `json:"text"`
}
type alarmMsg struct {
	Sensor string `json:"sensor_uuid"`
	SensorName string `json:"sensor_name"`
	Message string `json:"msg"`
}
func (r *alarmRing) exec(d *model.LogicData) {
	fmt.Printf("[alarm] %s\n", r.Message)

	r.ch <- alarmMsg{Sensor:d.SID, SensorName:d.SName, Message:r.Message}

	if r.next != nil {
		r.next.exec(d)
	}
}
