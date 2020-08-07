package main

import (
	"github.com/KumKeeHyun/PDK/application/dataService/sql"
	"github.com/KumKeeHyun/PDK/application/rest"
	"github.com/KumKeeHyun/PDK/application/usecase/nodeUsecase"
	"github.com/KumKeeHyun/PDK/application/usecase/sensorUsecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sql.Setup()

	nr := sql.NewNodeRepository()
	sr := sql.NewSensorRepository()

	nu := nodeUsecase.NewNodeUsecase(nr, sr)
	su := sensorUsecase.NewSensorUsecase(sr)

	h := rest.NewHandler(nu, su)

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

	r.GET("/registerInfo", h.RegisterInfo)

	r.Run()
}
