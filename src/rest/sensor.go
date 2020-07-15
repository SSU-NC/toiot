package rest

import (
	"pdk/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitSensorRoutes(e *gin.RouterGroup, h *controllers.Handler) {
	e.POST("/regist", h.AddSensor)

	e.GET("/info", h.GetSensorInfo)
}
