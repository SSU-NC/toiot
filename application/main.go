package main

import (
	"log"

	"github.com/KumKeeHyun/toiot/application/usecase/eventUsecase"

	"github.com/KumKeeHyun/toiot/application/dataService/sql"
	"github.com/KumKeeHyun/toiot/application/rest/handler"
	"github.com/KumKeeHyun/toiot/application/setting"
	"github.com/KumKeeHyun/toiot/application/usecase/registUsecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sql.Setup()

	sir := sql.NewSinkRepo()
	ndr := sql.NewNodeRepo()
	snr := sql.NewSensorRepo()
	lgr := sql.NewLogicRepo()
	lsr := sql.NewLogicServiceRepo()
	tpr := sql.NewTopicRepo()

	ru := registUsecase.NewRegistUsecase(sir, ndr, snr, lgr, lsr, tpr)
	eu := eventUsecase.NewEventUsecase(lsr)

	// TODO : init Topic table according to the setting value (example : Topic{"sensors", 3, 3})

	h := handler.NewHandler(ru, eu)

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	setRegistrationRoute(r, h)
	setEventRoute(r, h)

	log.Fatal(r.Run(setting.Appsetting.Server))
}

func setEventRoute(r *gin.Engine, h *handler.Handler) {
	event := r.Group("/event")
	{
		event.POST("", h.RegistLogicService)
	}
}

func setRegistrationRoute(r *gin.Engine, h *handler.Handler) {
	regist := r.Group("/regist")
	{
		sink := regist.Group("/sink")
		{
			sink.GET("", h.ListSinks)
			sink.POST("", h.RegistSink)
			sink.DELETE("/:id", h.UnregistSink)
		}
		node := regist.Group("/node")
		{
			node.GET("", h.ListNodes)
			node.POST("", h.RegistNode)
			node.DELETE("/:id", h.UnregistNode)
		}
		sensor := regist.Group("/sensor")
		{
			sensor.GET("", h.ListSensors)
			sensor.POST("", h.RegistSensor)
			sensor.DELETE("/:id", h.UnregistSensor)
		}
		logic := regist.Group("/logic")
		{
			logic.GET("", h.ListLogics)
			logic.POST("", h.RegistLogic)
			logic.DELETE("/:id", h.UnregistLogic)
		}
		logicService := regist.Group("/logic-service")
		{
			logicService.GET("", h.ListLogicServices)
			logicService.DELETE("/:id", h.UnregistLogicService)
		}
		topic := regist.Group("/topic")
		{
			topic.GET("", h.ListTopics)
			topic.POST("", h.RegistTopic)
			topic.DELETE("/:id", h.UnregistTopic)
		}
	}
}
