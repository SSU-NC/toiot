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
}

func (ls *Logic) Getenv() {
	ls.Server = os.Getenv("LOGIC_SERVER")
	if ls.Server == "" {
		ls.Server = "220.70.2.160:8082"
	}
}

var Logicsetting = &Logic{}

type Kafka struct {
	Broker      string   `toml:"broker"`
	GroupID     string   `toml:"group_id"`
	Topics      []string `toml:"topics"`
	ChanBufSize int      `toml:"chan_buf_size"`
}

func (ks *Kafka) Getenv() {
	GetenvStr(&ks.Broker, "220.70.2.1:9092", "KAFKA_BROKER")
	GetenvStr(&ks.GroupID, "logic", "KAFKA_GROUP")
	ks.Topics = []string{os.Getenv("KAFKA_TOPIC")}
	if ks.Topics[0] == "" {
		ks.Topics = []string{"sensors"}
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
		temp = "220.70.2.1:9200"
	}
	es.Addresses = []string{fmt.Sprintf("http://%s", temp)}
	GetenvInt(&es.RequestRetry, 3, "ELASTIC_RETRY")
	GetenvInt(&es.ChanBufSize, 500, "ELASTIC_BUFSIZE")
	GetenvInt(&es.BatchTicker, 5, "ELASTIC_BATCHTICKER")
	GetenvInt(&es.BatchSize, 450, "ELASTIC_BATCHSIZE")
}

var Elasticsetting = &Elastic{}

type MongoDB struct {
	Address string `toml:"address"`
	Port    string `toml:"port"`
}

func (ms *MongoDB) Getenv() {
	GetenvStr(&ms.Address, "127.0.0.1", "MONGO_ADDR")
	GetenvStr(&ms.Port, "27017", "MONGO_PORT")
}

var MongoDbSetting = &MongoDB{}

type App struct {
	Server string `toml:"address`
}

func (as *App) Getenv() {
	GetenvStr(&as.Server, "220.70.2.160:8081", "APP_SERVER")
}

var Appsetting = &App{}

func init() {
	Logicsetting.Getenv()
	Kafkasetting.Getenv()
	Elasticsetting.Getenv()
	Appsetting.Getenv()
	MongoDbSetting.Getenv()

	fmt.Println(Logicsetting, Kafkasetting, Elasticsetting, Appsetting)
}
