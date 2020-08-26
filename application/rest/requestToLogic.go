package rest

import (
	"time"

	"github.com/seheee/PDK/application/adapter"
	"github.com/seheee/PDK/application/domain/model"
	"github.com/go-resty/resty/v2"
)

func newNodeRequest(n adapter.Node) {
	node := struct {
		Name     string      `json:"name"`
		Group    string      `json:"group"`
		Location adapter.Loc `json:"location"`
	}{
		Name:     n.Name,
		Group:    n.Group,
		Location: n.Location,
	}

	cli := resty.New()
	resp, err := cli.SetTimeout(1 * time.Second).R().SetBody(node).Post("http://220.70.2.160:8090/metadata/node")
	if err != nil || resp.IsError() {
		//TODO
	}
}

func newSensorRequest(s model.Sensor) {
	cli := resty.New()
	resp, err := cli.SetTimeout(1 * time.Second).R().SetBody(s).Post("http://220.70.2.160:8090/metadata/sensor")
	if err != nil || resp.IsError() {
		//TODO
	}
}
