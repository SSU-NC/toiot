package logic

import (
	"fmt"
	"log"
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

type ActuatorElement struct {
	BaseElement
	
}
type Actuator struct {
	aid int		`json:"aid"`
	value int	`json:"value"`
	sleep int	`json:"sleep"`
}

func (ae *ActuatorElement) Exec(d *model.LogicData, addrs map[int]model.Sink) {
	/*
	Sinkaddr  돌면서 post요청 
	*/
	addrs:=[]string
	actuator:=Actuator{Value:10,sleep:10}
	//reqBody := sinkaddr
	pbytes, _ := json.Marshal(actuator)
	
	buff := bytes.NewBuffer(pbytes)
	
	for _, a := range addrs{
		resp, err := http.Post("http://"+a+"/act", "application/json", buff)
		if err != nil {
			panic(err)
		}
	}
	defer resp.Body.Close()
	// // Response 체크.
	// respBody, err := ioutil.ReadAll(resp.Body)
	// if err == nil {
	// 	str := string(respBody)
	// 	println(str)
	// }

}