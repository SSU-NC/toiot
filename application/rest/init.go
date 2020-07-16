package rest

import (
	"pdk/src/controllers"
	"pdk/src/setting"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunAPI() error {
	h, err := controllers.NewHandler()
	if err != nil {
		return err
	}
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	InitRoutes(r, h)

	return r.Run(setting.Serversetting.MakeAddr())
}

func InitRoutes(e *gin.Engine, h *controllers.Handler) {
	InitNodeRoutes(e.Group("/node"), h)
	InitSensorRoutes(e.Group("/sensor"), h)
}
