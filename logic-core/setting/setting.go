package setting

import (
	"fmt"
	"os"
	"strconv"
)

func GetenvInt(target *int, init int, env string) {
	var err error

	temp := os.Getenv(env)
	if temp == "" {
		*target = init
	} else {
		if *target, err = strconv.Atoi(temp); err != nil {
			*target = init
		}
	}
}

func GetenvStr(target *string, init, env string) {
	*target = os.Getenv(env)
	if *target == "" {
		*target = init
	}
}

type Logic struct {
	Server string
	Listen string
}

func (ls *Logic) Getenv() {
	GetenvStr(&ls.Server, "10.5.110.11:8084", "LOGIC_SERVER")
	GetenvStr(&ls.Listen, ls.Server, "LOGIC_LISTEN")
}

var Logicsetting = &Logic{}

type App struct {
	Server string
}

func (as *App) Getenv() {
	as.Server = os.Getenv("APP_SERVER")
	/*
		if as.Server == "" {
			as.Server = "localhost:8081"
		}
	*/
	if as.Server == "" {
		as.Server = "10.5.110.11:8081"
		// as.Server = "220.70.2.5:8081"
	}

}

var Appsetting = &App{}

type Kafka struct {
	Broker      string   `toml:"broker"`
	GroupID     string   `toml:"group_id"`
	Topics      []string `toml:"topics"`
	ChanBufSize int      `toml:"chan_buf_size"`
}

func (ks *Kafka) Getenv() {
	GetenvStr(&ks.Broker, "10.5.110.41:9092", "KAFKA_BROKER") //"localhost:9092", "KAFKA_BROKER")
	GetenvStr(&ks.GroupID, "logic1", "KAFKA_GROUP")
	ks.Topics = []string{os.Getenv("KAFKA_TOPIC")}
	if ks.Topics[0] == "" {
		ks.Topics = []string{"sensor-data"}
	}
	GetenvInt(&ks.ChanBufSize, 500, "KAFKA_BUFSIZE")
}

var Kafkasetting = &Kafka{}

type Elastic struct {
	Addresses    []string `toml:"addresses"`
	RequestRetry int      `toml:"request_retry"`
	ChanBufSize  int      `toml:"chan_buf_size"`
	BatchTicker  int      `toml:"batch_ticker"`
	BatchSize    int      `toml:"batch_size"`
}

func (es *Elastic) Getenv() {
	temp := os.Getenv("ELASTIC_SERVER")
	if temp == "" {
		temp = "10.5.110.38:9200" //"localhost:9200"
	}
	es.Addresses = []string{fmt.Sprintf("http://%s", temp)}
	GetenvInt(&es.RequestRetry, 3, "ELASTIC_RETRY")
	GetenvInt(&es.ChanBufSize, 500, "ELASTIC_BUFSIZE")
	GetenvInt(&es.BatchTicker, 5, "ELASTIC_BATCHTICKER")
	GetenvInt(&es.BatchSize, 450, "ELASTIC_BATCHSIZE")
}

var Elasticsetting = &Elastic{}

func init() {
	Logicsetting.Getenv()
	Kafkasetting.Getenv()
	Elasticsetting.Getenv()
	Appsetting.Getenv()

	fmt.Println(Logicsetting, Kafkasetting, Elasticsetting)
}
