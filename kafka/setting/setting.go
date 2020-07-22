package setting

import (
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
)

type Kafka struct {
	Broker  string   `toml:"broker"`
	GroupID string   `toml:"group_id"`
	Topics  []string `toml:"topics"`
}

var KafkaSetting = &Kafka{}

type Elastic struct {
	Addresses []string `toml:"addresses"`
}

var ElasticSetting = &Elastic{}

type WebSocket struct {
	URL string `toml:"url"`
}

var WebsocketSetting = &WebSocket{}

func Setup() {
	tree, err := toml.LoadFile("conf/config.toml")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf.config.toml': %v", err)
		return
	}

	kafkaTree := tree.Get("kafka").(*toml.Tree)
	kafkaTree.Unmarshal(KafkaSetting)

	elasticTree := tree.Get("elastic").(*toml.Tree)
	elasticTree.Unmarshal(ElasticSetting)

	wsTree := tree.Get("websocket").(*toml.Tree)
	wsTree.Unmarshal(WebsocketSetting)

	fmt.Println(KafkaSetting, ElasticSetting, WebsocketSetting)
}
