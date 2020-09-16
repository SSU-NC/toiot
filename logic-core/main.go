package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/trace"
	"syscall"
	"time"

	"github.com/KumKeeHyun/toiot/logic-core/usecase"

	"github.com/KumKeeHyun/toiot/logic-core/adapter"
	"github.com/KumKeeHyun/toiot/logic-core/dataService/memory"
	"github.com/KumKeeHyun/toiot/logic-core/elasticClient"
	"github.com/KumKeeHyun/toiot/logic-core/kafkaConsumer/sarama"
	"github.com/KumKeeHyun/toiot/logic-core/logicService"
	"github.com/KumKeeHyun/toiot/logic-core/rest/handler"
	"github.com/KumKeeHyun/toiot/logic-core/setting"
	"github.com/KumKeeHyun/toiot/logic-core/usecase/eventUC"
	"github.com/KumKeeHyun/toiot/logic-core/usecase/logicCoreUC"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// code for tracing goroutine
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("end")
		trace.Stop()
	}()

	rr := memory.NewRegistRepo()
	ks := sarama.NewKafkaConsumer()
	es := elasticClient.NewElasticClient()
	ls := logicService.NewLogicService()

	evuc := eventUC.NewEventUsecase(rr, ls)
	lcuc := logicCoreUC.NewLogicCoreUsecase(rr, ks, es, ls)

	h := handler.NewHandler(evuc, lcuc)
	r := gin.Default()
	SetEventRoute(r, h)
	RegistLogicService(evuc)

	go log.Fatal(r.Run(setting.Logicsetting.Server))

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
}

func SetEventRoute(r *gin.Engine, h *handler.Handler) {
	e := r.Group("/event")
	{
		e.POST("/sink/delete", h.DeleteSink)
		e.POST("/node/create", h.CreateNode)
		e.POST("/node/delete", h.DeleteNode)
		e.POST("/sensor/delete", h.DeleteSensor)
		e.POST("/logic/create", h.CreateLogic)
		e.POST("/logic/delete", h.DeleteLogic)
	}
}

func RegistLogicService(ls usecase.EventUsecase) {
	var (
		sinks  []adapter.Sink
		url    = fmt.Sprintf("http://%s/event", setting.Appsetting.Server)
		regist = adapter.LogicService{
			Addr: setting.Logicsetting.Listen,
			Topic: adapter.Topic{
				Name: setting.Kafkasetting.Topics[0],
			},
		}
	)

	client := resty.New()
	client.SetRetryCount(5).SetRetryWaitTime(10 * time.Second).SetRetryMaxWaitTime(30 * time.Second)
	resp, err := client.R().SetResult(&sinks).SetBody(regist).Post(url)
	if err != nil || !resp.IsSuccess() {
		panic(fmt.Errorf("can't regist logicService"))
	}

	for _, s := range sinks {
		for _, n := range s.Nodes {
			ls.CreateNode(&n, s.Name)
		}
	}
}
