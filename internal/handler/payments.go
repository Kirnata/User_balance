package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSinglePay(c *gin.Context) {
	//id, err := getUserId(c)
	//if err != nil {
	//	return
	//}

	//payId, err := h.service.Balance.GetBalance(id)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"balance": balance,
	//})
}

func (h *Handler) GetAllPayments(c *gin.Context) {

}

func (h *Handler) CreateB2BPay(c *gin.Context) {

}
