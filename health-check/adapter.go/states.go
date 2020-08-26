package adapter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/seheee/PDK/health-check/domain/model"
)

type States struct {
	Timestamp string      `json:"timestamp"`
	State     []NodeState `json:"state"`
}

type NodeState struct {
	NodeID string `json:"n_uuid"`
	State  bool   `json:"state"`
}

type HealthInfo struct {
	UUID  string `json:"n_uuid"`
	State int    `json:"state"`
}

func (s *States) GetNodeIDs() []string {
	if s == nil || s.State == nil || len(s.State) == 0 {
		return nil
	}
	res := make([]string, len(s.State))
	for i, v := range s.State {
		res[i] = v.NodeID
	}
	return res
}

func (s *States) GetNodeMap() map[string]NodeState {
	if s == nil || s.State == nil {
		return nil
	}
	res := map[string]NodeState{}
	for _, v := range s.State {
		res[v.NodeID] = v
	}
	return res
}

type StateDoc struct {
	Node      Node         `json:"node"`
	Status    model.Status `json:"status"`
	Timestamp string       `json:"timestamp"`
}

type Document struct {
	Index string
	Doc   StateDoc
}

func (d *Document) String() string {
	doc, err := json.Marshal(d.Doc)
	if err != nil {
		return ""
	}
	h := fmt.Sprintf(`{"index":{"_index":"%s"}}`, d.Index)
	return strings.Join([]string{h, "\n", string(doc), "\n"}, "")
}
