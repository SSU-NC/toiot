package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Document struct {
	Index string
	Doc   interface{}
}

func (d *Document) String() string {
	doc, err := json.Marshal(d.Doc)
	if err != nil {
		return ""
	}
	h := fmt.Sprintf("{\"index\":{\"_index\":\"%s\"}}\n", d.Index)
	return strings.Join([]string{h, string(doc), "\n"}, "")
}
