package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/trace"
	"syscall"

	"github.com/KumKeeHyun/toiot/logic-core/dataService/memory"
	"github.com/KumKeeHyun/toiot/logic-core/elasticClient"
	"github.com/KumKeeHyun/toiot/logic-core/kafkaConsumer/sarama"
	"github.com/KumKeeHyun/toiot/logic-core/logicService"
	"github.com/KumKeeHyun/toiot/logic-core/rest/handler"
	"github.com/KumKeeHyun/toiot/logic-core/setting"
	"github.com/KumKeeHyun/toiot/logic-core/usecase/eventUC"
	"github.com/KumKeeHyun/toiot/logic-core/usecase/logicCoreUC"
	"github.com/gin-gonic/gin"
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

	rr := memory.NewRegistRepo
	ks := sarama.NewKafkaConsumer()
	es := elasticClient.NewElasticClient()
	ls := logicService.NewLogicService()

	evuc := eventUC.NewEventUsecase(rr, ls)
	lcuc := logicCoreUC.NewLogicCoreUsecase(rr, ks, es, ls)

	h := handler.NewHandler(evuc, lcuc)
	r := gin.Default()
	SetEventRoute(r, h)

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
		e.POST("/sink/delete", h.DeleteSensor)
		e.POST("/logic/delete", h.CreateLogic)
		e.POST("/logic/delete", h.DeleteLogic)
	}
}
