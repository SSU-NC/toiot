package rest

import (
	"fmt"
	"time"

	"github.com/KumKeeHyun/PDK/application/setting"
	"github.com/go-resty/resty/v2"
)

func syncInfoRequest() {
	url := fmt.Sprintf("http://%s/syncInfo", setting.Logicsetting.Server)
	cli := resty.New()
	cli.SetTimeout(1 * time.Second).R().Post(url)
}
