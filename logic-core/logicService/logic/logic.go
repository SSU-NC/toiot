package logic

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KumKeeHyun/toiot/logic-core/domain/model"
)

type Elementer interface {
	SetNext(Elementer)
	Exec(*model.LogicData)
}

type BaseElement struct {
	next Elementer
}

func (e *BaseElement) SetNext(next Elementer) {
	e.next = next
}

func (e *BaseElement) Exec(d *model.LogicData) {
	if e.next != nil {
		e.next.Exec(d)
	}
}

func BuildLogic(l *model.Logic) (Elementer, error) {
	if len(l.Elems) == 0 {
		return nil, fmt.Errorf("invalid Element's length: %v", *l)
	}
	first, err := UnmarshalElement(&l.Elems[0])
	log.Println("in BuildLogic, l = ", l)
	if err != nil {
		return nil, err
	}
	res := &BaseElement{}
	res.SetNext(first)
	for _, raw := range l.Elems[1:] { // Elem 링크드 리스트 생성 후 리턴
		if elem, err := UnmarshalElement(&raw); err != nil {
			return nil, err
		} else {
			first.SetNext(elem)
			first = elem
		}
	}
	return res, nil
}

func UnmarshalElement(e *model.Element) (Elementer, error) {
	elem := GetElementer(e.Elem)
	log.Println("in UnmarshalElement, e(Element) = ", e)

	if elem == nil {
		return nil, fmt.Errorf("invalid Element : %s", e.Elem)
	}
	if bArg, err := json.Marshal(e.Arg); err == nil {
		log.Println("in UnmarshalElement, After Marshal bArg = ", string(bArg))
		if err = json.Unmarshal(bArg, elem); err != nil {
			return nil, err
		} else {
			log.Println("in UnmarshalElement, elem = ", elem)
			return elem, nil
		}
	} else {
		return nil, err
	}
}

func GetElementer(elem string) Elementer {
	switch elem {
	case "value":
		return &ValueElement{}
	case "time":
		return &TimeElement{}
	case "email":
		return &EmailElement{Interval: make(map[string]bool)}
	case "actuator":
		return &ActuatorElement{Interval: make(map[string]bool)}
	default:
		return nil
	}
}
