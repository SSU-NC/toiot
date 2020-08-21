package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/trace"
	"syscall"

	"github.com/KumKeeHyun/PDK/logic-core/dataService/memory"
	"github.com/KumKeeHyun/PDK/logic-core/elasticClient"
	kafkaConsumer "github.com/KumKeeHyun/PDK/logic-core/kafkaConsumer/sarama"
	"github.com/KumKeeHyun/PDK/logic-core/logicCore"
	"github.com/KumKeeHyun/PDK/logic-core/rest"
	_ "github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/KumKeeHyun/PDK/logic-core/usecase/logicCoreUC"
	"github.com/KumKeeHyun/PDK/logic-core/usecase/metaDataUC"
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
	ks := kafkaConsumer.NewKafkaConsumer()
	es := elasticClient.NewElasticClient()
	ls := logicCore.NewLogicCore()

	mduc := metaDataUC.NewMetaDataUsecase(mr, ls)
	lcuc := logicCoreUC.NewLogicCoreUsecase(mr, ks, es, ls)

	h := rest.NewHandler(mduc, lcuc)
	go rest.RunServer(h)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
}
