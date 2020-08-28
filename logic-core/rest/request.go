package rest

import (
	"errors"
	"fmt"

	"github.com/KumKeeHyun/PDK/logic-core/adapter"
	"github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/go-resty/resty/v2"
)

func MetaInfoRequest() (adapter.MetaInfo, error) {
	res := adapter.MetaInfo{}

	url := fmt.Sprintf("http://%s/registerInfo", setting.Appsetting.Server)
	client := resty.New()
	resp, err := client.R().SetResult(&res).Get(url)
	if err != nil || resp.IsError() {
		return adapter.MetaInfo{}, errors.New("registerInfo request error")
	}

	return res, nil
}
