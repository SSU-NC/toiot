package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/KumKeeHyun/toiot/application/adapter"
	"github.com/KumKeeHyun/toiot/application/domain/model"
	"github.com/gin-gonic/gin"
)

// ListSinks ...
// @Summary List sink node(raspi info)
// @Description get sinks list
// @Tags sink
// @Param  page query int false "page num"
// @Param  size query int false "page size(row)"
// @Produce  json
// @Success 200 {array} model.Sink "default, return all sinks."
// @Success 201 {object} adapter.SinkPage "if page query is exist, return pagenation result. pages only valid when page is 1."
// @Router /regist/sink [get]
func (h *Handler) ListSinks(c *gin.Context) {
	var (
		err   error
		sinks []model.Sink
		page  adapter.Page
		pages int
	)
	if c.Bind(&page); page.IsBinded() {
		if page.Size == 0 {
			page.Size = 10
		}
		if sinks, err = h.ru.GetSinksPage(page); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if page.Page == 1 {
			pages = h.ru.GetSinkPageCount(page.Size)
		}
		c.JSON(http.StatusOK, gin.H{"sinks": sinks, "pages": pages})
		return
	} else {
		sinks, err := h.ru.GetSinks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, sinks)
		return
	}
}

// RegistSink ...
// @Summary Add sink node(raspi info)
// @Description Add sink node
// @Tags sink
// @Param  sink body model.Sink true "name, address(only ip address, don't include port)"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Sink "include topic info"
// @Router /regist/sink [post]
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
	h.eu.CreateSinkEvent(&sink)
	c.JSON(http.StatusOK, sink)
}

// UnregistSink ...
// @Summary Delete sink node(raspi info)
// @Description Delete sink node
// @Tags sink
// @Param  id path int true "sink's id"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Sink "include topic, nodes info"
// @Router /regist/sink [delete]
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

// ListNodes ...
// @Summary List sensor node
// @Description get nodes listh.eu.CreateNodeEvent(&node)
// @Tags node
// @Param  page query int false "page num"
// @Param  size query int false "page size(row)"
// @Param  sink query int false "sink filter"
// @Param  left query float32 false "location(longitude) filter"
// @Param  right query float32 false "location(longitude) filter"
// @Param  up query float32 false "location(Latitude) filter"
// @Param  down query float32 false "location(Latitude) filter"
// @Produce  json
// @Success 200 {array} model.Node "default, return all nodes. if location query is exist, return location filter result(square)."
// @Success 201 {object} adapter.NodePage "if page query is exist, return pagenation result. pages only valid when page is 1."
// @Router /regist/node [get]
func (h *Handler) ListNodes(c *gin.Context) {
	var (
		err    error
		nodes  []model.Node
		page   adapter.Page
		pages  int
		square adapter.Square
	)

	if c.Bind(&page); page.IsBinded() {
		if page.Size == 0 {
			page.Size = 10
		}
		if nodes, err = h.ru.GetNodesPage(page); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if page.Page == 1 {
			pages = h.ru.GetNodePageCount(page)
		}
		c.JSON(http.StatusOK, gin.H{"nodes": nodes, "pages": pages})
		return
	} else if c.Bind((&square)); square.IsBinded() {
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

// RegistNode ...
// @Summary Add sensor node
// @Description Add sensor node
// @Tags node
// @Param  node body model.Node true "name, lat, lng, sink_id"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Node "include sink, sink.topic, sensors, sensors.logics info"
// @Router /regist/node [post]
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
	go h.eu.PostToSink(node.SinkID)
	c.JSON(http.StatusOK, node)

}

// UnregistNode ...
// @Summary Delete sensor node
// @Description Delete sensor node
// @Tags node
// @Param  id path int true "node's id"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Node "include sink, sink.topic info"
// @Router /regist/node [delete]
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
	go h.eu.PostToSink(node.SinkID)
	c.JSON(http.StatusOK, node)
}

// ListSensors ...
// @Summary List sensor info
// @Description get sensors list
// @Tags sensor
// @Param  page query int false "page num"
// @Param  size query int false "page size(row)"
// @Produce  json
// @Success 200 {array} model.Sensor "default, return all sensors."
// @Success 201 {object} adapter.SensorPage "if page query is exist, return pagenation result. pages only valid when page is 1."
// @Router /regist/sensor [get]
func (h *Handler) ListSensors(c *gin.Context) {
	var (
		err     error
		sensors []model.Sensor
		page    adapter.Page
		pages   int
	)

	if c.Bind(&page); page.IsBinded() {
		if page.Size == 0 {
			page.Size = 10
		}
		if sensors, err = h.ru.GetSensorsPage(page); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if page.Page == 1 {
			pages = h.ru.GetSensorPageCount(page.Size)
		}
		c.JSON(http.StatusOK, gin.H{"sensors": sensors, "pages": pages})
		return
	} else {
		sensors, err := h.ru.GetSensors()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, sensors)
		return
	}

}

// RegistSensor ...
// @Summary Add sensor info
// @Description Add sensor info
// @Tags sensor
// @Param  sensor body model.Sensor true "name, sensorValues(only value name)"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Node "include sensorValues info"
// @Router /regist/sensor [post]
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

// UnregistSensor ...
// @Summary Delete sensor
// @Description Delete sensor
// @Tags sensor
// @Param  id path int true "sensor's id"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Sensor "include logics info"
// @Router /regist/sensor [delete]
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

// ListLogics ...
// @Summary List logics info
// @Description get logics list
// @Tags logic
// @Produce  json
// @Success 200 {array} model.Logic "return all logics info."
// @Router /regist/logic [get]
func (h *Handler) ListLogics(c *gin.Context) {
	logics, err := h.ru.GetLogics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	aLogics := adapter.LogicsToAdapter(logics)
	c.JSON(http.StatusOK, aLogics)
}

// ListActuator ...
func (h *Handler) ListActuators(c *gin.Context) {
	var (
		err       error
		actuators []model.Actuator
		page      adapter.Page
		pages     int
	)
	log.Println("in ListActuators")
	if c.Bind(&page); page.IsBinded() {
		if page.Size == 0 {
			page.Size = 10
		}
		if actuators, err = h.ru.GetActuatorsPage(page); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if page.Page == 1 {
			pages = h.ru.GetActuatorPageCount(page.Size)
		}
		c.JSON(http.StatusOK, gin.H{"actuators": actuators, "pages": pages})
		return
	} else {
		actuators, err := h.ru.GetActuators()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, actuators)
		return
	}

}

// RegistActuator ...
func (h *Handler) RegistActuator(c *gin.Context) {
	var actuator model.Actuator
	log.Println("in RegistActuator")
	if err := c.ShouldBindJSON(&actuator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.ru.RegistActuator(&actuator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actuator)
}

// UnregistActuator ...
func (h *Handler) UnregistActuator(c *gin.Context) {
	log.Println("in UnregistActuator")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actuator := model.Actuator{ID: id}

	err = h.ru.UnregistActuator(&actuator)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, actuator)
}

// RegistLogic ...
// @Summary Add logic info
// @Description Add logic info
// @Tags logic
// @Param  logic body adapter.Logic true "logic_name, elems"
// @Accept  json
// @Produce  json
// @Success 200 {object} adapter.Logic "include sensor info"
// @Router /regist/logic [post]
func (h *Handler) RegistLogic(c *gin.Context) {
	var aLogic adapter.Logic
	if err := c.ShouldBindJSON(&aLogic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("aLogic = ", aLogic)
	logic, err := adapter.LogicToModel(&aLogic)
	log.Println("logic = ", logic)
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

// UnregistLogic ...
// @Summary Delete logic
// @Description Delete logic
// @Tags logic
// @Param  id path int true "logic's id"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Logic "include sensor info"
// @Router /regist/logic [delete]
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

// ListLogicServices ...
// @Summary List LogicServices info
// @Description get LogicServices list
// @Tags LogicService
// @Produce  json
// @Success 200 {array} model.LogicService "return all logics info."
// @Router /regist/logic-service [get]
func (h *Handler) ListLogicServices(c *gin.Context) {
	logicServices, err := h.ru.GetLogicServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logicServices)
}

// UnregistLogicService ...
// @Summary Delete LogicService
// @Description Delete LogicService
// @Tags logicService
// @Param  id path int true "logicSerivce's id"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Logic "include topic info"
// @Router /regist/logic-service [delete]
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

// ListTopics ...
// @Summary List topics info
// @Description get topics list
// @Tags topic
// @Produce  json
// @Success 200 {array} model.Topic "return all topics info."
// @Router /regist/topic [get]
func (h *Handler) ListTopics(c *gin.Context) {
	topics, err := h.ru.GetTopics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, topics)
}

// RegistTopic ...
// @Summary Add topic info
// @Description Add topic info
// @Tags topic
// @Param  logic body model.Logic true "name, partitions, replications"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Topic
// @Router /regist/topic [post]
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

// UnregistTopic ...
// @Summary Delete topic(kafka topic for logicservices)
// @Description Delete topic(kafka topic for logicservices)
// @Tags topic
// @Param  id path int true "topic's id"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Topic "include logicService info"
// @Router /regist/topic [delete]
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
