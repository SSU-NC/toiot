package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"

	"github.com/KumKeeHyun/toiot/logic-core/adapter"
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
	log.Println("in Email's Exec")
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

type ActuatorElement struct {
	BaseElement
	Aid      int      `json:"aid"`
	Values   []Values `json:"values"`
	Sleep    int      `json:"sleep"`
	Interval map[string]bool
}

type Actuator struct {
	Nid    int      `json;"nid"`
	Aid    int      `json:"aid"`
	Values []Values `json:"values"`
}
type Values struct {
	Value int `json:"value"`
	Sleep int `json:"sleep"`
}

func (ae *ActuatorElement) Exec(d *model.LogicData) {
	/*
		Sinkaddr  돌면서 post요청
	*/
	//

	ok, exist := ae.Interval[d.Node.Name]
	if !exist {
		ae.Interval[d.Node.Name] = true
	}
	if ok {
		ae.Interval[d.Node.Name] = false

		res := Actuator{
			Nid:    d.Node.Nid,
			Aid:    ae.Aid,
			Values: ae.Values,
		}
		pbytes, _ := json.Marshal(res)
		buff := bytes.NewBuffer(pbytes)
		addr := (*adapter.AddrMap)[d.Node.Sid]
		log.Println("FLAG : logic-action - ActuatorElement - Exec")
		resp, err := http.Post("http://"+addr.Addr+"/actuator", "application/json", buff)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		tick := time.Tick(3 * time.Minute)
		go func() {
			<-tick
			ae.Interval[d.Node.Name] = true
		}()
	}
	ae.BaseElement.Exec(d)
	// // Response 체크.
	// respBody, err := ioutil.ReadAll(resp.Body)
	// if err == nil {
	// 	str := string(respBody)
	// 	println(str)
	// }

}
