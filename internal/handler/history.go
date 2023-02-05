package handler

import (
	"github.com/Kirnata/User_balance"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type getHistoryResponse struct {
	Data []User_balance.History `json:"data"`
}

func (h *Handler) GetHistory(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	log.Println("hhuihuihiu")
	history, err := h.service.History.GetHistory(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getHistoryResponse{
		Data: history,
	})
}
