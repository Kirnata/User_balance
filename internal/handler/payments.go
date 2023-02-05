package handler

import (
	"github.com/Kirnata/User_balance"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateSinglePay(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input User_balance.SinglePay
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	payId, err := h.service.Payments.CreateSinglePay(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"payId": payId,
	})
}

func (h *Handler) CreateB2BPay(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input User_balance.B2BPay
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	payId, err := h.service.Payments.CreateB2BPay(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"payId": payId,
	})
}
