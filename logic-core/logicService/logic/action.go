package logic

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
)

const (
	from    = "toiotpdk@gmail.com"
	pass    = "ndsprnlulncwgdvo"
	bodyFmt = "sensor(%s) on node(%s)"
	msgFmt  = "Form: %s\nTo: %s\nSubject: ToIoT email\n\n%s"
)

type EmailElement struct {
	BaseElement
	Email    string `json:"text"`
	Interval map[string]bool
}

func (ee *EmailElement) Exec(d *model.LogicData) {
	ok, exist := ee.Interval[d.Node.Name]
	if !exist {
		ee.Interval[d.Node.Name] = true
	}
	if ok {
		ee.Interval[d.Node.Name] = false

		body := fmt.Sprintf(bodyFmt, d.SensorName, d.Node.Name)
		msg := fmt.Sprintf(msgFmt, from, ee.Email, body)

		smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{ee.Email}, []byte(msg))

		tick := time.Tick(3 * time.Minute)
		go func() {
			<-tick
			ee.Interval[d.Node.Name] = true
		}()
	}

	ee.BaseElement.Exec(d)
}
