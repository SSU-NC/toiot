package rest

import (
	"github.com/seheee/PDK/logic-core/setting"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func RunServer(h *Handler) error {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.POST("/metadata/node", h.NewNode)
	r.POST("/metadata/sensor", h.NewSensor)
	r.DELETE("/metadata/node", h.DeleteNode)
	r.DELETE("/metadata/sensor", h.DeleteSensor)

	r.POST("/logiccore", h.NewLogicChain)
	r.GET("/logiccore", h.GetAllLogic)
	r.DELETE("/logiccore", h.DeleteLogicChain)

	r.GET("/websocket", h.NewWebSocket)

	return r.Run(setting.Serversetting.MakeAddr())
}
