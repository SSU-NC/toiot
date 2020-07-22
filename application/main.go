package main

import (
	"github.com/KumKeeHyun/PDK/application/interface/db/orm"
	"github.com/KumKeeHyun/PDK/application/interface/handler"
	"github.com/KumKeeHyun/PDK/application/service/kafka"
	"github.com/KumKeeHyun/PDK/application/setting"
	"github.com/KumKeeHyun/PDK/application/usecase/nodeUsecase"
	"github.com/KumKeeHyun/PDK/application/usecase/sensorUsecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	setting.Setup()
	orm.Setup()
	kafka.Setup()

	nr := orm.NewNodeRepository()
	sr := orm.NewSensorRepository()

	nu := nodeUsecase.NewNodeUsecase(nr, sr)
	su := sensorUsecase.NewSensorUsecase(sr)

	h := handler.NewHandler(nu, su)

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	ng := r.Group("/node")
	{
		ng.GET("", h.GetAllInfo)
		ng.POST("", h.RegisterNode)
		ng.DELETE("", h.DeleteNode)
	}
	sg := r.Group("/sensor")
	{
		sg.GET("", h.GetSensorsInfo)
		sg.POST("", h.RegisterSensor)
		sg.DELETE("", h.DeleteSensor)
	}

	r.GET("/kafkaConsumerManager", h.KafkaConsumerManager)

	r.Run()
}
