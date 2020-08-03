/*
TODO : ring(node) chain
*/

package logicCore

import (
	"fmt"

	"github.com/KumKeeHyun/PDK/logic-core/domain/model"
)

type elasticRing struct {
	chainRingBase
}

func (r *elasticRing) execute(d *model.LogicData) {
	fmt.Printf("tempElasticRing: %v\n", *d)
}

type locFilterRing struct {
	chainRingBase
	Loc map[string]bool `json:"loc"`
}

func (r *locFilterRing) execute(d *model.LogicData) {
	if _, ok := r.Loc[d.NodeInfo.Group]; ok {
		d.SName += "[loc]"
		r.chainRingBase.execute(d)
	}
}

type rangeRing struct {
	chainRingBase
	Value string  `json:"value"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
}

func (r *rangeRing) execute(d *model.LogicData) {
	isRange := func(val float64) bool {
		if val >= r.Min && val < r.Max {
			return true
		} else {
			return false
		}
	}
	v, ok := d.Values[r.Value]
	if !ok {
		return
	}
	d.SName += "[range]"
	if isRange(v) {
		r.chainRingBase.execute(d)
	}
}
