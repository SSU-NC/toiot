package rest

import (
	"github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	r.POST("/syncInfo", h.SyncMetaInfo)

	r.POST("/logiccore", h.NewLogicChain)
	r.GET("/logiccore", h.GetAllLogic)
	r.DELETE("/logiccore", h.DeleteLogicChain)

	r.GET("/websocket", h.NewWebSocket)
	return r.Run(setting.Logicsetting.Server)
}
