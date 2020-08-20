package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/KumKeeHyun/PDK/health-check/dataService/memory"
	"github.com/KumKeeHyun/PDK/health-check/elasticClient"
	"github.com/KumKeeHyun/PDK/health-check/kafkaConsumer"
	"github.com/KumKeeHyun/PDK/health-check/usecase/eventUC"
	"github.com/KumKeeHyun/PDK/health-check/usecase/statusCheckUC"
)

func main() {
	sr := memory.NewStatusRepo()
	ks := kafkaConsumer.NewKafkaConsumer()
	es := elasticClient.NewElasticClient()

	_ = statusCheckUC.NewStatusCheckUsecase(sr)
	_ = eventUC.NewEventUsecase(sr, ks, es)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
}
