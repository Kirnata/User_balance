package handler

import (
	"github.com/Kirnata/User_balance/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", h.signUp)
		auth.POST("/sign_in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.POST("/b2b_pay", h.CreateB2BPay)
		api.POST("/single_pay", h.CreateSinglePay)
		api.GET("/get_history", h.GetHistory)
		api.GET("/get_balance", h.getBalance)
	}
	return router
}
