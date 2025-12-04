package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ispu-monitoring/backend/internal/model"
	"github.com/ispu-monitoring/backend/internal/service"
)

type DashboardHandler struct {
	service *service.DashboardService
}

func NewDashboardHandler(service *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

// GetOverview handles GET /api/v1/dashboard/overview
func (h *DashboardHandler) GetOverview(c *gin.Context) {
	overview, err := h.service.GetOverview()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "FETCH_ERROR",
				Message: "Failed to fetch dashboard overview",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Dashboard overview retrieved successfully",
		Data:    overview,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// GetStatistics handles GET /api/v1/dashboard/statistics
func (h *DashboardHandler) GetStatistics(c *gin.Context) {
	overview, err := h.service.GetOverview()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "FETCH_ERROR",
				Message: "Failed to fetch statistics",
				Details: err.Error(),
			},
		})
		return
	}
	
	statistics := gin.H{
		"summary":                overview.Summary,
		"category_distribution":  overview.CategoryDistribution,
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Statistics retrieved successfully",
		Data:    statistics,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// GetCategories handles GET /api/v1/categories
func (h *DashboardHandler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "FETCH_ERROR",
				Message: "Failed to fetch categories",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Categories retrieved successfully",
		Data:    categories,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}
