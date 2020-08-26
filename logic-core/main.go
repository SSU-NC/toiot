package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/trace"
	"syscall"

	"github.com/seheee/PDK/logic-core/dataService/memory"
	"github.com/seheee/PDK/logic-core/elasticClient"
	kafkaConsumer "github.com/seheee/PDK/logic-core/kafkaConsumer/sarama"
	"github.com/seheee/PDK/logic-core/logicCore"
	"github.com/seheee/PDK/logic-core/rest"
	_ "github.com/seheee/PDK/logic-core/setting"
	"github.com/seheee/PDK/logic-core/usecase/logicCoreUC"
	"github.com/seheee/PDK/logic-core/usecase/metaDataUC"
	"github.com/seheee/PDK/logic-core/usecase/websocketUC"
	"github.com/seheee/PDK/logic-core/db"

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

	mr := memory.NewMetaRepo()
	lr := db.NewLogicRepository()
	ks := kafkaConsumer.NewKafkaConsumer()
	es := elasticClient.NewElasticClient()
	ls := logicCore.NewLogicCore()

	event := make(chan interface{}, 2)
	mduc := metaDataUC.NewMetaDataUsecase(mr, ls)
	lcuc := logicCoreUC.NewLogicCoreUsecase(mr, lr, ks, es, ls, event)
	wuc := websocketUC.NewWebsocketUsecase(event)
	
	h := rest.NewHandler(mduc, lcuc, wuc)
	go rest.RunServer(h)
	
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
}
