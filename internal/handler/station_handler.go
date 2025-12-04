package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ispu-monitoring/backend/internal/model"
	"github.com/ispu-monitoring/backend/internal/service"
)

type StationHandler struct {
	service          *service.StationService
	dashboardService *service.DashboardService
}

func NewStationHandler(service *service.StationService, dashboardService *service.DashboardService) *StationHandler {
	return &StationHandler{
		service:          service,
		dashboardService: dashboardService,
	}
}

// GetAllStations handles GET /api/v1/stations
func (h *StationHandler) GetAllStations(c *gin.Context) {
	province := c.Query("province")
	
	var stations []model.Station
	var err error
	
	if province != "" {
		stations, err = h.service.GetStationsByProvince(province)
	} else {
		stations, err = h.service.GetAllStations()
	}
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "FETCH_ERROR",
				Message: "Failed to fetch stations",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Stations retrieved successfully",
		Data:    stations,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// GetStationByID handles GET /api/v1/stations/:id
func (h *StationHandler) GetStationByID(c *gin.Context) {
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
	
	station, err := h.service.GetStationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "NOT_FOUND",
				Message: "Station not found",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Station retrieved successfully",
		Data:    station,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// GetStationLatestData handles GET /api/v1/stations/:id/latest
func (h *StationHandler) GetStationLatestData(c *gin.Context) {
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
	
	station, err := h.service.GetStationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "NOT_FOUND",
				Message: "Station not found",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Station data retrieved successfully",
		Data:    station,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// CreateStation handles POST /api/v1/stations
func (h *StationHandler) CreateStation(c *gin.Context) {
	var station model.Station
	
	if err := c.ShouldBindJSON(&station); err != nil {
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
	
	if err := h.service.CreateStation(&station); err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "CREATE_ERROR",
				Message: "Failed to create station",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusCreated, model.APIResponse{
		Success: true,
		Message: "Station created successfully",
		Data:    station,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// UpdateStation handles PUT /api/v1/stations/:id
func (h *StationHandler) UpdateStation(c *gin.Context) {
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
	
	var station model.Station
	if err := c.ShouldBindJSON(&station); err != nil {
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
	
	if err := h.service.UpdateStation(uint(id), &station); err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "UPDATE_ERROR",
				Message: "Failed to update station",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Station updated successfully",
		Data:    station,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// DeleteStation handles DELETE /api/v1/stations/:id
func (h *StationHandler) DeleteStation(c *gin.Context) {
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
	
	if err := h.service.DeleteStation(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "DELETE_ERROR",
				Message: "Failed to delete station",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Station deleted successfully",
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}

// GetMapStations handles GET /api/v1/map/stations
func (h *StationHandler) GetMapStations(c *gin.Context) {
	// Use dashboard service to get map stations with ISPU data
	mapStations, err := h.dashboardService.GetMapStationsData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Error: &model.APIError{
				Code:    "FETCH_ERROR",
				Message: "Failed to fetch map stations",
				Details: err.Error(),
			},
		})
		return
	}
	
	c.JSON(http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Map stations retrieved successfully",
		Data:    mapStations,
		Meta: &model.MetaData{
			Timestamp: time.Now(),
			Version:   "1.0.0",
		},
	})
}
