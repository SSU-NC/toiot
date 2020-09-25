package healthCheckUC

import (
	"fmt"
	"time"

	"github.com/KumKeeHyun/toiot/health-check/adapter"
	"github.com/KumKeeHyun/toiot/health-check/setting"
	"github.com/go-resty/resty/v2"
)

var (
	appClient *resty.Client
	url       string
)

func init() {
	appClient = resty.New()
	appClient.SetRetryCount(2).SetRetryWaitTime(100 * time.Millisecond).SetTimeout(500 * time.Millisecond)
	url = fmt.Sprintf("http://%s%s", setting.Appsetting.Server, setting.Appsetting.RequestPath)
}

func getSinkList() ([]adapter.Sink, error) {
	res := []adapter.Sink{}
	resp, err := appClient.R().SetResult(&res).Get(url)
	if resp.IsError() {
		return nil, err
	}
	return res, nil
}
