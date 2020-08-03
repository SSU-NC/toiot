package logicCore

import "github.com/KumKeeHyun/PDK/logic-core/domain/model"

func getRinger(logic string) chainRing {
	// TODO

	switch logic {
	case "range":
		return &rangeRing{}
	case "loc":
		return &locFilterRing{}
	case "elastic":
		return &elasticRing{}
	default:
		return nil
	}
}

func chainFactory(rs []model.LogicRing) chainRing {
	// TODO

	return &chainRingBase{}
}
