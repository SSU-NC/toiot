package controllers

import (
	"net/http"
	"pdk/src/models"

	"github.com/gin-gonic/gin"
)

// GetRegistInfo : GET /node/regist
func (h *Handler) GetRegistInfo(c *gin.Context) {
	nodes, err := h.db.GetAllNodes()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := range nodes {
		nodes[i].Sensors, err = h.db.GetSensorsByNID(nodes[i].UUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for j := range nodes[i].Sensors {
			nodes[i].Sensors[j].ValueNames, err = h.db.GetSensorValues(nodes[i].Sensors[j].UUID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
	}
	c.JSON(http.StatusOK, nodes)
}

// AddNode : POST /node/regist -d {"name" : "nv", "location" : "lv"}
func (h *Handler) AddNode(c *gin.Context) {
	var node = models.NewNode()
	err := c.ShouldBindJSON(&node)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node, err = h.db.AddNode(node)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if node.Sensors != nil {
		for _, v := range node.Sensors {
			ns := models.NodeSensor{
				NodeUUID:   node.UUID,
				SensorUUID: v.UUID,
			}
			h.db.AddNodeSensor(ns)
		}
	}
	c.JSON(http.StatusOK, node)
}

// AddNodeSensor : POST /node/sensor -d {"node_id" : 1, "sensor_id" : 2}
func (h *Handler) AddNodeSensor(c *gin.Context) {
	var ns = models.NodeSensor{}
	err := c.ShouldBindJSON(&ns)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ns, err = h.db.AddNodeSensor(ns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ns)
}
