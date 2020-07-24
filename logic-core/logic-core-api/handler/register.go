package handler

import (
	"net/http"

	"github.com/KumKeeHyun/PDK/logic-core/logic-core-api/model"
	"github.com/gin-gonic/gin"
)

func NewNode(c *gin.Context) {
	var node model.AppNode
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n := model.ToNode(node)
	model.RegisterRepo.AddNode(node.UUID, n)
	c.JSON(http.StatusOK, node)
}

func DeleteNode(c *gin.Context) {
	var node model.AppNode
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.RegisterRepo.DelNode(node.UUID)
	c.JSON(http.StatusOK, node)
}

func NewSensor(c *gin.Context) {
	var sensor model.AppSensor
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := model.ToSensor(sensor)
	model.RegisterRepo.AddSensor(sensor.UUID, s)
	c.JSON(http.StatusOK, sensor)
}

func DeleteSensor(c *gin.Context) {
	var sensor model.AppSensor
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	model.RegisterRepo.DelSensor(sensor.UUID)
	c.JSON(http.StatusOK, sensor)
}
