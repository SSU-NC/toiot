package rest

import (
	"github.com/KumKeeHyun/PDK/logic-core/setting"
	"github.com/gin-gonic/gin"
)

func RunServer(h *Handler) error {
	r := gin.Default()

	r.POST("/metadata/node", h.NewNode)
	r.POST("/metadata/sensor", h.NewSensor)
	r.DELETE("/metadata/node", h.DeleteNode)
	r.DELETE("/metadata/sensor", h.DeleteSensor)
	r.POST("/syncInfo", h.SyncMetaInfo)

	return r.Run(setting.Logicsetting.Server)
}
