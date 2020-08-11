/*
TODO : ring(node) chain
*/

package logicCore

import (
	"fmt"
	"github.com/seheee/PDK/logic-core/domain/model"
)

type Ringer interface {
	setNext(Ringer)
	exec(*model.LogicData)
}

type baseRing struct {
	next Ringer
}

func (r *baseRing) setNext(n Ringer) {
	r.next = n
}
func (r *baseRing) exec(d *model.LogicData) {
	fmt.Printf("test: %v\n", *d)
	if r.next != nil {
		r.next.exec(d)
	}
}
/*
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
}*/
