package memory

import (
	"github.com/KumKeeHyun/PDK/logic-core/adapter"
	"github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/go-resty/resty/v2"
)

func initMetaRepoRequest(mr *metaRepo) {
	if mr == nil {
		return
	}

	temp := struct {
		NInfo []adapter.Node   `json:"node_info"`
		SInfo []adapter.Sensor `json:"sensor_info"`
	}{}

	cli := resty.New()
	if _, err := cli.R().SetResult(&temp).Get(setting.AppServerSetting.Address + "/registerInfo"); err != nil {
		return
	}

	for _, n := range temp.NInfo {
		mn := adapter.AppToNode(&n)
		mr.NewNode(n.UUID, &mn)
	}
	for _, s := range temp.SInfo {
		ms := adapter.AppToSensor(&s)
		mr.NewSensor(s.UUID, &ms)
	}
}
