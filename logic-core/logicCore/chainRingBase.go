/*
TODO : ring(node) chain
*/

package logicCore

import (
	"fmt"

	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
)

type chainRing interface {
	setNext(nr chainRing)
	execute(d *model.LogicData)
}

type chainRingBase struct {
	next chainRing
}

func (crb *chainRingBase) setNext(nr chainRing) {
	crb.next = nr
}

func (crb *chainRingBase) execute(d *model.LogicData) {
	fmt.Printf("test: %v\n", *d)
	if crb.next != nil {
		crb.next.execute(d)
	}
}
