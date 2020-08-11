package logicCore

import (
	"fmt"
	"encoding/json"
	"github.com/seheee/PDK/logic-core/domain/model"
)

/*type RingRequest struct {
	Sensor string `json:"sensor"`
	Logics []struct {
		Logic string `json:"logic"`
		Arg map[string]interface{} `json:"arg"`
	} `json:"logics"`
}*/

func getRinger(logic string) Ringer {
	switch logic {
	case "range":
		return &rangeRing{}
	case "group":
		return &groupRing{}
	case "time":
		return &timeRing{}
	case "email":
		return &emailRing{}
	case "alarm":
		return &alarmRing{}
	default:
		return nil
	}
}

func UnmarshalRing(l string, a interface{}) Ringer {
	var res Ringer
	if res = getRinger(l); res == nil {
		fmt.Println("ring not exist")
		return nil
	}
	if arg, err := json.Marshal(a); err == nil {
		if err = json.Unmarshal(arg, res); err == nil {
			return res
		} else {
			fmt.Println("json unmarshal error")
			fmt.Println("error", err.Error())
			return nil
		}
	} else {
		fmt.Println("json marshal error")
		return nil
	}
}

func chainFactory(rr *model.RingRequest) *baseRing {

	var chain, res Ringer
	if res = UnmarshalRing(rr.Logic[0].Elem, rr.Logic[0].Arg); res == nil {
		fmt.Printf("cannot unmarshal ring %s\n", rr.Logic[0].Elem)
		return nil
	}
	var base baseRing
	base.setNext(res)
	for _, logic := range rr.Logic[1:] {
		/*if logic.Elem == "empty" {
			continue
		}*/
		if chain = UnmarshalRing(logic.Elem, logic.Arg); chain == nil {
			fmt.Printf("cannot unmarshal ring %s\n", logic.Elem)
			return nil
		}
		fmt.Println(chain)
		res.setNext(chain)
		res = chain
	}

	//in := make(chan Meta, 100)
	//Mux.AddChan(rr.Sensor, in)
	return &base
}
