package rest

import (
	"net/http"
	"os"
	"os/exec"
	"pdk/src/controllers"
	"pdk/src/setting"

	"github.com/gin-gonic/gin"
)

func RunAPI() error {
	h, err := controllers.NewHandler()
	if err != nil {
		return err
	}
	r := gin.Default()
	InitRoutes(r, h)

	return r.Run(setting.Serversetting.MakeAddr())
}

func InitRoutes(e *gin.Engine, h *controllers.Handler) {
	InitNodeRoutes(e.Group("/node"), h)
	InitSensorRoutes(e.Group("/sensor"), h)
	e.GET("/test", func(c *gin.Context) {
		cmd := exec.Command("kafka/kafkaConsum")
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "none"})
		}

	})
}
