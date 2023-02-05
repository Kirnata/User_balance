package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getBalance(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	balance, err := h.service.Balance.GetBalance(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}
