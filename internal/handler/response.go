package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	zap.Error(errors.New(message))
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
