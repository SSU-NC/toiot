package setting

import (
	"os"
	"strconv"
)

type Status struct {
	Count int
	Tick  int
	Drop  int
}

func (ss *Status) Getenv() {
	var err error

	cnt := os.Getenv("STATUS_COUNT")
	if cnt == "" {
		ss.Count = 4
	} else {
		if ss.Count, err = strconv.Atoi(cnt); err != nil {
			ss.Count = 4
		}
	}

	tick := os.Getenv("STATUS_TICK")
	if cnt == "" {
		ss.Tick = 5
	} else {
		if ss.Tick, err = strconv.Atoi(tick); err != nil {
			ss.Tick = 5
		}
	}

	drop := os.Getenv("STATUS_DROP")
	if cnt == "" {
		ss.Drop = 12
	} else {
		if ss.Drop, err = strconv.Atoi(drop); err != nil {
			ss.Drop = 5
		}
	}
}

var StatusSetting = &Status{}

type Kafka struct {
	Broker      string   `toml:"broker"`
	GroupID     string   `toml:"group_id"`
	Topics      []string `toml:"topics"`
	ChanBufSize int      `toml:"chan_buf_size"`
}

func (ks *Kafka) Getenv() {
	var err error

	ks.Broker = os.Getenv("KAFKA_BROKER")
	if ks.Broker == "" {
		ks.Broker = "220.70.2.1:9092"
	}

	ks.GroupID = os.Getenv("KAFKA_GROUP")
	if ks.GroupID == "" {
		ks.GroupID = "hc"
	}

	ks.Topics = []string{os.Getenv("KAFKA_TOPIC")}
	if ks.Topics[0] == "" {
		ks.Topics = []string{"healthcheck"}
	}

	bufSize := os.Getenv("KAFKA_BUFSIZE")
	if bufSize == "" {
		ks.ChanBufSize = 10
	} else {
		if ks.ChanBufSize, err = strconv.Atoi(bufSize); err != nil {
			ks.ChanBufSize = 10
		}
	}
}

var KafkaSetting = &Kafka{}

type Elastic struct {
	Addresses []string `toml:"addresses"`
}

func (es *Elastic) Getenv() {
	es.Addresses = []string{os.Getenv("ELASTIC_SERVER")}
	if es.Addresses[0] == "" {
		es.Addresses = []string{"http://220.70.2.1:9200/"}
	}
}

var ElasticSetting = &Elastic{}

func init() {
	StatusSetting.Getenv()
	KafkaSetting.Getenv()
}
