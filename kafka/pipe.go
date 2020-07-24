package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/KumKeeHyun/PDK/kafka/elasticPipe"
	"github.com/KumKeeHyun/PDK/kafka/kafkaPipe"
	"github.com/KumKeeHyun/PDK/kafka/processing"
	"github.com/KumKeeHyun/PDK/kafka/setting"
	"github.com/KumKeeHyun/PDK/kafka/wsClient"
)

func main() {
	setting.Setup()

	sock := wsClient.SetupAndStart()
	defer sock.Close()

	if err := kafkaPipe.Setup(); err != nil {
		log.Fatal(err.Error())
		return
	}
	if _, err := elasticPipe.Setup(); err != nil {
		log.Fatal(err.Error())
		return
	}

	end := make(chan os.Signal)
	signal.Notify(end, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-end; os.Exit(0) }()

	done := elasticPipe.PushToElastic(processing.ProcessingPipe(kafkaPipe.StartConsumer()))
	for res := range done {
		fmt.Println(res)
	}
}
