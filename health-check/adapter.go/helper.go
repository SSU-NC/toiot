package adapter

import (
	"errors"
	"fmt"

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

	cli := resty.New()
	_, err := cli.R().SetResult(&res).SetQueryString(queryFac(ids)).Get("http://220.70.2.160:8080/node/select")
	if err != nil {
		return nil, err
	}

	return res, nil
}
