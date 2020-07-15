package controllers

import (
	"net/http"
	"pdk/src/models"

	"github.com/gin-gonic/gin"
)

// GetAllSensorInfo : GET
func (h *Handler) GetSensorInfo(c *gin.Context) {
	ss, err := h.db.GetAllSensors()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := range ss {
		ss[i].ValueNames, err = h.db.GetSensorValues(ss[i].UUID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, ss)
}

// AddSensor : POST -d {"name" : "nv", "num_of_values" : 2, "value_names" : ["v1", "v2"]}
func (h *Handler) AddSensor(c *gin.Context) {
	var sensor = models.NewSensor()
	err := c.ShouldBindJSON(&sensor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sensor, err = h.db.AddSensor(sensor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i, v := range sensor.ValueNames {
		sv := models.SensorValue{
			SensorUUID: sensor.UUID,
			ValueName:  v,
			Index:      i,
		}
		h.db.AddSensorValue(sv)
	}
	c.JSON(http.StatusOK, sensor)
}
