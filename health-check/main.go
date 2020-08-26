package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/seheee/PDK/health-check/setting"
	"github.com/seheee/PDK/health-check/usecase/websocketUC"

	"github.com/seheee/PDK/health-check/dataService/memory"
	"github.com/seheee/PDK/health-check/elasticClient"
	"github.com/seheee/PDK/health-check/kafkaConsumer"
	"github.com/seheee/PDK/health-check/usecase/eventUC"
	"github.com/seheee/PDK/health-check/usecase/statusCheckUC"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	sr := memory.NewStatusRepo()
	ks := kafkaConsumer.NewKafkaConsumer()
	es := elasticClient.NewElasticClient()

	event := make(chan struct{}, 2)
	_ = statusCheckUC.NewStatusCheckUsecase(sr, event)
	_ = eventUC.NewEventUsecase(sr, ks, es, event)
	wu := websocketUC.NewWebsocketUsecase(sr, event)

	r := gin.New()

	r.GET("/health-check", func(c *gin.Context) {
		listen := make(chan interface{})
		wu.Register(listen)
		defer wu.Unregister(listen)

		conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
		if err != nil {
			log.Printf("upgrade: %s", err.Error())
		}
		fmt.Println("connect websocket!")

		conn.WriteJSON(sr.GetHealthInfo())

		for data := range listen {
			conn.WriteJSON(data)
		}
		fmt.Println("disconnect websocket!")
	})

	go log.Fatal(r.Run(setting.Appsetting.Server))

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
}
