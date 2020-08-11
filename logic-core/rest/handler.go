package rest

import (
	"net/http"

	"github.com/seheee/PDK/logic-core/domain/model"
	"github.com/seheee/PDK/logic-core/adapter"
	"github.com/seheee/PDK/logic-core/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	mduc usecase.MetaDataUsecase
	lcuc usecase.LogicCoreUsecase
}


func NewHandler(mduc usecase.MetaDataUsecase, lcuc usecase.LogicCoreUsecase) *Handler {
	return &Handler{
		mduc: mduc,
		lcuc: lcuc,
	}
}

func (h *Handler) NewLogicChain(c *gin.Context) {
	var rr model.RingRequest
	if err := c.ShouldBindJSON(&rr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.lcuc.SetLogicChain(&rr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, rr)
	}
}

func (h *Handler) DeleteLogicChain(c *gin.Context) {
	var rr model.RingRequest
	if err := c.ShouldBindJSON(&rr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.lcuc.RemoveLogicChain(rr.LogicName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, rr)
	}
}

func (h *Handler) DeleteLogicChain(c *gin.Context) {
	var rr model.RingRequest
	if err := c.ShouldBindJSON(&rr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.lcuc.RemoveLogicChainBySID(rr.Sensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, rr)
	}
}

func (h *Handler) NewNode(c *gin.Context) {
	var an adapter.Node
	if err := c.ShouldBindJSON(&an); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n := adapter.AppToNode(&an)

	if res, err := h.mduc.NewNode(an.UUID, &n); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (h *Handler) NewSensor(c *gin.Context) {
	var as adapter.Sensor
	if err := c.ShouldBindJSON(&as); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := adapter.AppToSensor(&as)

	if res, err := h.mduc.NewSensor(as.UUID, &s); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (h *Handler) DeleteNode(c *gin.Context) {
	var an adapter.Node
	if err := c.ShouldBindJSON(&an); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.mduc.DeleteNode(an.UUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, an)
	}
}

func (h *Handler) DeleteSensor(c *gin.Context) {
	var as adapter.Sensor
	if err := c.ShouldBindJSON(&as); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.mduc.DeleteSensor(as.UUID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, as)
	}
}
