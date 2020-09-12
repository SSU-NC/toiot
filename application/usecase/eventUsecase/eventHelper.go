package eventUsecase

import (
	"fmt"
	"time"

	"github.com/KumKeeHyun/toiot/application/domain/model"
	"github.com/go-resty/resty/v2"
)

var (
	pingClient  *resty.Client
	eventClient *resty.Client
)

func init() {
	eventClient = resty.New()
	eventClient.SetRetryCount(3).SetRetryWaitTime(300 * time.Millisecond).SetRetryMaxWaitTime(1 * time.Second)
	pingClient = resty.New()
	pingClient.SetTimeout(200 * time.Millisecond)
}

func ping(l model.LogicService) error {
	path := "/ping"
	url := makeUrl(l.Addr, path)

	resp, _ := pingClient.R().Get(url)
	if resp.IsSuccess() {
		return nil
	}
	return fmt.Errorf("%s response error : %d", l.Addr, resp.StatusCode())
}

func makeUrl(addr, path string) string {
	return fmt.Sprintf("http://%s%s", addr, path)
}
