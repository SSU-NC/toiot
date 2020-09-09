package setting

import (
	"fmt"
	"os"
	"strconv"
)

type Health struct {
	Server string
}

func (hs *Health) Getenv() {
	hs.Server = os.Getenv("HEALTH_SERVER")
	if hs.Server == "" {
		hs.Server = "0.0.0.0:8083"
	}
}

var Healthsetting = &Health{}

type App struct {
	Server      string
	MetaRequest string
}

func (as *App) Getenv() {
	as.Server = os.Getenv("APP_SERVER")
	if as.Server == "" {
		as.Server = "0.0.0.0:8081"
	}
	as.MetaRequest = "/node/select"
}

var Appsetting = &App{}

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
		ss.Tick = 30
	} else {
		if ss.Tick, err = strconv.Atoi(tick); err != nil {
			ss.Tick = 30
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
		ks.GroupID = "health-check"
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

func init() {
	Healthsetting.Getenv()
	Appsetting.Getenv()
	StatusSetting.Getenv()
	KafkaSetting.Getenv()

	fmt.Printf("Health : &v\nApp : %v\nStatus : %v\nKafka : %v\n\n", Healthsetting, Appsetting, StatusSetting, KafkaSetting)
}
