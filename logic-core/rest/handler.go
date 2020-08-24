package rest

import (
	"net/http"
	"fmt"
	"github.com/seheee/PDK/logic-core/domain/model"
	"github.com/seheee/PDK/logic-core/adapter"
	"github.com/seheee/PDK/logic-core/usecase"
	//"github.com/seheee/PDK/logic-core/logicCore"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	mduc usecase.MetaDataUsecase
	lcuc usecase.LogicCoreUsecase
	wuc usecase.WebsocketUsecase
}


func NewHandler(mduc usecase.MetaDataUsecase, lcuc usecase.LogicCoreUsecase, wuc usecase.WebsocketUsecase) *Handler {
	return &Handler{
		mduc: mduc,
		lcuc: lcuc,
		wuc: wuc,
	}
}

func (h *Handler) NewWebSocket(c *gin.Context) {
	listen := make(chan interface{})
	h.wuc.Register(listen)
	defer h.wuc.Unregister(listen)

	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		fmt.Printf("upgrade: %s", err.Error())
	}
	fmt.Println("connect websocket!")

	for data := range listen {
		conn.WriteJSON(data)
	}
	fmt.Println("disconnect websocket!")
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
	var r struct{Id string `json:"id"`}
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.lcuc.RemoveLogicChain(r.Id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, r)
	}
}

func (h *Handler) DeleteLogicChains(c *gin.Context) {
	var rr model.RingRequest
	if err := c.ShouldBindJSON(&rr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.lcuc.RemoveLogicChainsBySID(rr.Sensor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, rr)
	}
}

func (h *Handler) GetAllLogic(c *gin.Context) {
	logics, err := h.lcuc.GetAllLogics()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logics)
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
