package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KumKeeHyun/toiot/application/adapter"
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"github.com/gin-gonic/gin"
)

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

	h.eu.DeleteSinkEvent(&sink)
	c.JSON(http.StatusOK, sink)
}

func (h *Handler) ListNodes(c *gin.Context) {
	var (
		err    error
		nodes  []model.Node
		page   adapter.Page
		pages  int
		square adapter.Square
	)

	if temp := c.Query("page"); temp != "" {
		c.ShouldBind(&page)
		fmt.Println(page)
		if page.Size == 0 {
			page.Size = 10
		}
		if nodes, err = h.ru.GetNodesPage(page); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if page.Page == 1 {
			pages = h.ru.GetPageCount(page.Size)
		}
		c.JSON(http.StatusOK, gin.H{"nodes": nodes, "pages": pages})
		return
	} else if temp := c.Query("left"); temp != "" {
		c.ShouldBind(&square)
		if nodes, err = h.ru.GetNodesSquare(square); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nodes)
		return
	} else {
		nodes, err := h.ru.GetNodes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nodes)
		return
	}

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

	h.eu.CreateNodeEvent(&node)
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

	h.eu.DeleteNodeEvent(&node)
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

	err := h.ru.RegistSensor(&sensor)
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

	h.eu.DeleteSensorEvent(&sensor)
	c.JSON(http.StatusOK, sensor)
}

func (h *Handler) ListLogics(c *gin.Context) {
	logics, err := h.ru.GetLogics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	aLogics := adapter.LogicsToAdapter(logics)
	c.JSON(http.StatusOK, aLogics)
}

func (h *Handler) RegistLogic(c *gin.Context) {
	var aLogic adapter.Logic
	if err := c.ShouldBindJSON(&aLogic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logic, err := adapter.LogicToModel(&aLogic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.ru.RegistLogic(&logic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resLogic, _ := adapter.LogicToAdapter(&logic)
	h.eu.CreateLogicEvent(&logic)
	c.JSON(http.StatusOK, resLogic)
}

func (h *Handler) UnregistLogic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logic := model.Logic{ID: id}

	err = h.ru.UnregistLogic(&logic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resLogic, _ := adapter.LogicToAdapter(&logic)
	h.eu.DeleteLogicEvent(&logic)
	c.JSON(http.StatusOK, resLogic)
}

func (h *Handler) ListLogicServices(c *gin.Context) {
	logicServices, err := h.ru.GetLogicServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logicServices)
}

func (h *Handler) UnregistLogicService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logicService := model.LogicService{ID: id}

	err = h.ru.UnregistLogicService(&logicService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logicService)
}

func (h *Handler) ListTopics(c *gin.Context) {
	topics, err := h.ru.GetTopics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topics)
}

func (h *Handler) RegistTopic(c *gin.Context) {
	var topic model.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.ru.RegistTopic(&topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topic)
}

func (h *Handler) UnregistTopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	topic := model.Topic{ID: id}

	err = h.ru.UnregistTopic(&topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topic)
}
