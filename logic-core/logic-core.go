package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/KumKeeHyun/PDK/logic-core/core/pipe/elasticPipe"
	"github.com/KumKeeHyun/PDK/logic-core/core/pipe/kafkaPipe"
	"github.com/KumKeeHyun/PDK/logic-core/core/pipe/processing"
	"github.com/KumKeeHyun/PDK/logic-core/logic-core-api/handler"
	"github.com/KumKeeHyun/PDK/logic-core/logic-core-api/model"
	"github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/gin-gonic/gin"
)

func main() {
	setting.Setup()

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
	go func() {
		for res := range done {
			fmt.Println(res)
		}
	}()

	resp, err := http.Get("http://220.70.2.160:8080/registerInfo")
	if err == nil && resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			var jsonBody model.InitRegister
			if err = json.Unmarshal(body, &jsonBody); err == nil {
				for _, appn := range jsonBody.NodeInfo {
					n := model.ToNode(appn)
					model.RegisterRepo.AddNode(appn.UUID, n)
				}
				for _, apps := range jsonBody.SensorInfo {
					s := model.ToSensor(apps)
					model.RegisterRepo.AddSensor(apps.UUID, s)
				}
			}
		}
		fmt.Println(model.RegisterRepo)
		resp.Body.Close()
	}

	r := gin.Default()
	r.POST("/register/newnode", handler.NewNode)
	r.POST("/register/deletenode", handler.DeleteNode)
	r.POST("/register/newsensor", handler.NewSensor)
	r.POST("/register/deletesensor", handler.DeleteSensor)
	r.Run(setting.Serversetting.MakeAddr())
}
