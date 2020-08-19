package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KumKeeHyun/PDK/application/adapter"
	"github.com/KumKeeHyun/PDK/application/domain/model"
	"github.com/KumKeeHyun/PDK/application/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	siu usecase.SinkUsecase
	nu  usecase.NodeUsecase
	su  usecase.SensorUsecase
}

func NewHandler(siu usecase.SinkUsecase, nu usecase.NodeUsecase, su usecase.SensorUsecase) *Handler {
	return &Handler{
		siu: siu,
		nu:  nu,
		su:  su,
	}
}

func (h *Handler) GetSinkInfo(c *gin.Context) {
	sinks, err := h.siu.GetAllSinks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sinks)
}

func (h *Handler) GetSinkByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sink, err := h.siu.GetSinkByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *sink)
}

func (h *Handler) RegisterSink(c *gin.Context) {
	var sink model.Sink
	if err := c.ShouldBindJSON(&sink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sink.CheckIP(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	new, err := h.siu.RegisterSink(&sink)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *new)
}

func (h *Handler) DeleteSink(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sink := model.Sink{
		ID: uint(id),
	}

	err = h.siu.DeleteSink(&sink)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sink)
}

func (h *Handler) GetNodesInfo(c *gin.Context) {
	nodes, err := h.nu.GetAllNodesWithSensorsWithValues()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ns := adapter.ModelsToNodes(nodes)
	c.JSON(http.StatusOK, ns)
}

func (h *Handler) GetNodesByIDs(c *gin.Context) {
	ids := c.QueryArray("uuid")
	nodes, err := h.nu.GetNodesByUUID(ids)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ns := adapter.ModelsToNodes(nodes)
	c.JSON(http.StatusOK, ns)
}

func (h *Handler) RegisterNode(c *gin.Context) {
	var node adapter.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	n := adapter.NodeToModel(&node)
	new, err := h.nu.RegisterNode(&n)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newNodeRequest(node)
	c.JSON(http.StatusOK, *new)

}

func (h *Handler) DeleteNode(c *gin.Context) {
	var node adapter.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	n := adapter.NodeToModel(&node)

	dn, err := h.nu.DeleteNode(&n)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *dn)
}

func (h *Handler) GetSensorsInfo(c *gin.Context) {
	sensors, err := h.su.GetAllSensorsWithValues()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sensors)
}

func (h *Handler) RegisterSensor(c *gin.Context) {
	var sensor model.Sensor

	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new, err := h.su.RegisterSensor(&sensor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newSensorRequest(sensor)
	c.JSON(http.StatusOK, *new)
}

func (h *Handler) DeleteSensor(c *gin.Context) {
	var sensor model.Sensor
	if err := c.ShouldBindJSON(&sensor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ds, err := h.su.DeleteSensor(&sensor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *ds)
}

func (h *Handler) RegisterInfo(c *gin.Context) {
	nodeInfo, _ := h.nu.GetAllNodes()
	sensorInfo, _ := h.su.GetAllSensorsWithValues()
	msg := map[string]interface{}{
		"node_info":   nodeInfo,
		"sensor_info": sensorInfo,
	}
	c.JSON(http.StatusOK, msg)
}

func (h *Handler) CreateLogic(c *gin.Context) {
	var test struct {
		SensorUUID string                   `json:"sensor_uuid"`
		LogicName  string                   `json:"logic_name"`
		Logics     []map[string]interface{} `json:"logic"`
	}

	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("[new logic chain]\n%v\n", test)
	c.JSON(http.StatusOK, test)
}
