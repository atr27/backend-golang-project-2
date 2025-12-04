package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ispu-monitoring/backend/internal/model"
)

// HealthCheck handles health check endpoint
func HealthCheck(c *gin.Context) {
	response := model.APIResponse{
		Success: true,
		Message: "ISPU Monitoring API is running",
		Data: gin.H{
			"status":    "healthy",
			"timestamp": time.Now(),
			"version":   "1.0.0",
		},
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	}
	c.JSON(http.StatusOK, response)
}
