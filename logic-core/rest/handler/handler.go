package handler

import (
	"log"
	"net/http"

	"github.com/KumKeeHyun/toiot/logic-core/adapter"
	"github.com/KumKeeHyun/toiot/logic-core/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	evuc usecase.EventUsecase
	lcuc usecase.LogicCoreUsecase
}

func NewHandler(evuc usecase.EventUsecase, lcuc usecase.LogicCoreUsecase) *Handler {
	return &Handler{
		evuc: evuc,
		lcuc: lcuc,
	}
}

func (h *Handler) CreateSink(c *gin.Context) {
	var addr adapter.SinkAddr
	log.Println("in CreateSink")
	if err := c.ShouldBindJSON(&addr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.lcuc.AppendSinkAddr(&addr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, addr)

	// var an adapter.Node
	// if err := c.ShouldBindJSON(&an); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// if err := h.evuc.CreateNode(&an, an.Sink.Name); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, an)
}

func (h *Handler) DeleteSink(c *gin.Context) {
	var nl []adapter.Node
	if err := c.ShouldBindJSON(&nl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.evuc.DeleteSink(nl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nl)
}

func (h *Handler) CreateNode(c *gin.Context) {
	var an adapter.Node
	if err := c.ShouldBindJSON(&an); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.evuc.CreateNode(&an, an.Sink.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, an)
}

func (h *Handler) DeleteNode(c *gin.Context) {
	var an adapter.Node
	if err := c.ShouldBindJSON(&an); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.evuc.DeleteNode(&an); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, an)
}

func (h *Handler) DeleteSensor(c *gin.Context) {
	var as adapter.Sensor
	if err := c.ShouldBindJSON(&as); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.evuc.DeleteSensor(&as); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, as)
}

func (h *Handler) CreateLogic(c *gin.Context) {
	var al adapter.Logic
	log.Println("in createLogic")
	if err := c.ShouldBindJSON(&al); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.evuc.CreateLogic(&al); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, al)
}

func (h *Handler) DeleteLogic(c *gin.Context) {
	var al adapter.Logic
	if err := c.ShouldBindJSON(&al); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.evuc.DeleteLogic(&al); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, al)
}
