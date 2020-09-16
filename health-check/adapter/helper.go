package adapter

import (
	"errors"
	"fmt"

	"github.com/KumKeeHyun/PDK/health-check/setting"

	"github.com/go-resty/resty/v2"
)

func GetNodeInfo(s *States) ([]Node, error) {
	queryFac := func(vs []string) string {
		key := "uuid"
		query := fmt.Sprintf("%s=%s", key, vs[0])
		for _, v := range vs[1:] {
			query += fmt.Sprintf("&%s=%s", key, v)
		}
		return query
	}

	ids := s.GetNodeIDs()
	if ids == nil {
		return nil, errors.New("state : node_id does not exist")
	}
	res := make([]Node, 0)

	url := fmt.Sprintf("http://%s%s", setting.Appsetting.Server, setting.Appsetting.MetaRequest)
	cli := resty.New()
	_, err := cli.R().SetResult(&res).SetQueryString(queryFac(ids)).Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
