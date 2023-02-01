package handler

import (
	"github.com/Kirnata/User_balance/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.POST("/", h.CreateTransaction)
		api.GET("/", h.GetAllTransaction)
	}
	return router
}