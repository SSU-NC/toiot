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
	eventClient.SetTimeout(200 * time.Millisecond)
	pingClient = resty.New()
	pingClient.SetRetryCount(2).SetRetryWaitTime(100 * time.Millisecond).SetTimeout(500 * time.Millisecond)
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
