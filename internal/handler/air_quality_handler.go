package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ispu-monitoring/backend/internal/model"
	"github.com/ispu-monitoring/backend/internal/service"
)

type AirQualityHandler struct {
	service *service.AirQualityService
}

func NewAirQualityHandler(service *service.AirQualityService) *AirQualityHandler {
	return &AirQualityHandler{service: service}
}

// GetLatestData handles GET /api/v1/air-quality/latest
func (h *AirQualityHandler) GetLatestData(c *gin.Context) {
	data, err := h.service.GetLatestData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "FETCH_ERROR",
				Message: "Failed to fetch latest air quality data",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Latest air quality data retrieved successfully",
		Data:    data,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// GetStationHistory handles GET /api/v1/air-quality/station/:id
func (h *AirQualityHandler) GetStationHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "INVALID_ID",
				Message: "Invalid station ID",
				Details: err.Error(),
			},
		})
		return
	}
	
	// Parse query parameters for date range
	startDateStr := c.DefaultQuery("start_date", time.Now().AddDate(0, 0, -7).Format("2006-01-02"))
	endDateStr := c.DefaultQuery("end_date", time.Now().Format("2006-01-02"))
	
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "INVALID_DATE",
				Message: "Invalid start date format. Use YYYY-MM-DD",
				Details: err.Error(),
			},
		})
		return
	}
	
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "INVALID_DATE",
				Message: "Invalid end date format. Use YYYY-MM-DD",
				Details: err.Error(),
			},
		})
		return
	}
	
	// Add 23:59:59 to end date to include the entire day
	endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	
	history, err := h.service.GetHistoricalData(uint(id), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "FETCH_ERROR",
				Message: "Failed to fetch historical data",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Historical air quality data retrieved successfully",
		Data:    history,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// InsertAirQuality handles POST /api/v1/air-quality
func (h *AirQualityHandler) InsertAirQuality(c *gin.Context) {
	var data model.AirQuality
	
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "VALIDATION_ERROR",
				Message: "Invalid request data",
				Details: err.Error(),
			},
		})
		return
	}
	
	// Set timestamp if not provided
	if data.Timestamp.IsZero() {
		data.Timestamp = time.Now()
	}
	
	if err := h.service.InsertAirQuality(&data); err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "INSERT_ERROR",
				Message: "Failed to insert air quality data",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusCreated, model.APIResponse{
		Success: true,
		Message: "Air quality data inserted successfully",
		Data:    data,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}
