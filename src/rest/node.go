package rest

import (
	"pdk/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitNodeRoutes(e *gin.RouterGroup, h *controllers.Handler) {
	e.POST("/regist", h.AddNode)
	e.POST("/sensor", h.AddNodeSensor)

	e.GET("/regist", h.GetRegistInfo)
}
