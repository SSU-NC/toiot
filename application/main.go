package main

import (
	"log"

	"github.com/KumKeeHyun/PDK/application/dataService/sql"
	"github.com/KumKeeHyun/PDK/application/rest"
	"github.com/KumKeeHyun/PDK/application/setting"
	"github.com/KumKeeHyun/PDK/application/usecase/nodeUsecase"
	"github.com/KumKeeHyun/PDK/application/usecase/sensorUsecase"
	"github.com/KumKeeHyun/PDK/application/usecase/sinkUsecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sql.Setup()

	sir := sql.NewSinkRepository()
	nr := sql.NewNodeRepository()
	sr := sql.NewSensorRepository()

	siu := sinkUsecase.NewSinkUsecase(sir, nr)
	nu := nodeUsecase.NewNodeUsecase(nr, sr)
	su := sensorUsecase.NewSensorUsecase(sr)

	h := rest.NewHandler(siu, nu, su)

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	sig := r.Group("/sink")
	{
		sig.GET("", h.GetSinkInfo)
		sig.GET("/:id", h.GetSinkByID)
		sig.POST("", h.RegisterSink)
		sig.DELETE("/:id", h.DeleteSink)
	}
	ng := r.Group("/node")
	{
		ng.GET("", h.GetNodesInfo)
		ng.GET("/select", h.GetNodesByIDs)
		ng.POST("", h.RegisterNode)
		ng.DELETE("", h.DeleteNode)
	}
	sg := r.Group("/sensor")
	{
		sg.GET("", h.GetSensorsInfo)
		sg.POST("", h.RegisterSensor)
		sg.DELETE("", h.DeleteSensor)
	}

	r.GET("/registerInfo", h.RegisterInfo)

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, setting.Appsetting.React)
	})

	log.Fatal(r.Run(setting.Appsetting.Server))
}
