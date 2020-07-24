package setting

import (
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
)

type Server struct {
	Address string
	Port    string
}

func (c *Server) MakeAddr() string {
	return fmt.Sprintf("%s:%s", c.Address, c.Port)
}

var Serversetting = &Server{}

type Kafka struct {
	Broker         string   `toml:"broker"`
	GroupID        string   `toml:"group_id"`
	Topics         []string `toml:"topics"`
	NumOfConsumers int      `toml:"num_of_consumers"`
}

var KafkaSetting = &Kafka{}

type Elastic struct {
	Addresses []string `toml:"addresses"`
}

var ElasticSetting = &Elastic{}

func Setup() {
	tree, err := toml.LoadFile("conf/config.toml")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf.config.toml': %v", err)
		return
	}

	serverTree := tree.Get("server").(*toml.Tree)
	serverTree.Unmarshal(Serversetting)

	kafkaTree := tree.Get("kafka").(*toml.Tree)
	kafkaTree.Unmarshal(KafkaSetting)

	elasticTree := tree.Get("elastic").(*toml.Tree)
	elasticTree.Unmarshal(ElasticSetting)

	fmt.Println(Serversetting, KafkaSetting, ElasticSetting)
}
