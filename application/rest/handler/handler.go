package rest

import (
	"net/http"
	"strconv"

	"github.com/KumKeeHyun/toiot/application/domain/model"
	"github.com/KumKeeHyun/toiot/application/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ru usecase.RegistUsecase
}

func NewHandler(ru usecase.RegistUsecase) *Handler {
	return &Handler{
		ru: ru,
	}
}

func (h *Handler) ListSinks(c *gin.Context) {
	sinks, err := h.ru.GetSinks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sinks)
}

func (h *Handler) RegistSink(c *gin.Context) {
	var sink model.Sink
	if err := c.ShouldBindJSON(&sink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.ru.RegistSink(&sink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sink)
}

func (h *Handler) UnregistSink(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sink := model.Sink{ID: id}
	err = h.ru.UnregistSink(&sink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sink)
}

func (h *Handler) ListNodes(c *gin.Context) {
	nodes, err := h.ru.GetNodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

func (h *Handler) RegistNode(c *gin.Context) {
	var node model.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.ru.RegistNode(&node)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)

}

func (h *Handler) UnregistNode(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node := model.Node{ID: id}

	err = h.ru.UnregistNode(&node)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)
}

func (h *Handler) ListSensors(c *gin.Context) {
	sensors, err := h.ru.GetSensors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensors)
}

func (h *Handler) RegistSensor(c *gin.Context) {
	var sensor model.Sensor
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.ru.UnregistSensor(&sensor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensor)
}

func (h *Handler) UnregistSensor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sensor := model.Sensor{ID: id}

	err = h.ru.UnregistSensor(&sensor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensor)
}
