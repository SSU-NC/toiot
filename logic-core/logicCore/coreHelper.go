package logicCore

import (
	"fmt"
	"encoding/json"

	"github.com/seheee/PDK/logic-core/domain/model"
)

func getRinger(logic string, event chan interface{}) Ringer {
	switch logic {
	case "value":
		return &rangeRing{}
	case "group":
		return &groupRing{}
	case "time":
		return &timeRing{}
	case "email":
		return &emailRing{Time:true}
	case "alarm":
		return &alarmRing{ch: event}
	default:
		return nil
	}
}

func UnmarshalRing(l string, a interface{}, event chan interface{}) Ringer {
	var res Ringer
	if res = getRinger(l, event); res == nil {
		fmt.Println(l)
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

func chainFactory(rr *model.RingRequest, event chan interface{}) *baseRing {

	var chain, res Ringer
	if res = UnmarshalRing(rr.Logic[0].Elem, rr.Logic[0].Arg, event); res == nil {
		fmt.Printf("cannot unmarshal ring %s\n", rr.Logic[0].Elem)
		return nil
	}
	var base baseRing
	base.setNext(res)
	for _, logic := range rr.Logic[1:] {
		if chain = UnmarshalRing(logic.Elem, logic.Arg, event); chain == nil {
			fmt.Printf("cannot unmarshal ring %s\n", logic.Elem)
			return nil
		}
		res.setNext(chain)
		fmt.Printf("%+v\n", chain)
		res = chain
	}

	return &base
}
