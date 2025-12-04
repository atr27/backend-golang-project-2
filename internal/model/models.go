package model

import (
	"time"
)

// Station represents a monitoring station
type Station struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null" binding:"required"`
	Code      string    `json:"code" gorm:"uniqueIndex;not null" binding:"required"`
	Type      string    `json:"type" gorm:"not null" binding:"required"` // KLHK/INTEGRASI
	Latitude  float64   `json:"latitude" gorm:"not null" binding:"required"`
	Longitude float64   `json:"longitude" gorm:"not null" binding:"required"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AirQuality represents air quality measurement
type AirQuality struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	StationID uint       `json:"station_id" gorm:"index;not null"`
	Station   *Station   `json:"station,omitempty" gorm:"foreignKey:StationID"`
	ISPU      int        `json:"ispu" gorm:"not null" binding:"required,min=0"`
	PM25      *float64   `json:"pm25"`
	PM10      *float64   `json:"pm10"`
	CO        *float64   `json:"co"`
	NO2       *float64   `json:"no2"`
	O3        *float64   `json:"o3"`
	SO2       *float64   `json:"so2"`
	HC        *float64   `json:"hc"`
	Timestamp time.Time  `json:"timestamp" gorm:"index;not null"`
	Category  string     `json:"category" gorm:"-"`
	Color     string     `json:"color" gorm:"-"`
	CreatedAt time.Time  `json:"created_at"`
}

// ISPUCategory represents air quality category
type ISPUCategory struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	MinValue    int       `json:"min_value" gorm:"not null"`
	MaxValue    *int      `json:"max_value"`
	Category    string    `json:"category" gorm:"not null"`
	Description string    `json:"description"`
	Color       string    `json:"color" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
}

// DashboardOverview represents dashboard summary data
type DashboardOverview struct {
	Summary              DashboardSummary             `json:"summary"`
	CategoryDistribution map[string]int               `json:"category_distribution"`
	RecentReadings       []StationWithAirQuality      `json:"recent_readings"`
	ProvinceStats        []ProvinceStatistic          `json:"province_stats"`
}

// DashboardSummary represents summary statistics
type DashboardSummary struct {
	TotalStations  int64     `json:"total_stations"`
	ActiveStations int64     `json:"active_stations"`
	LastUpdate     time.Time `json:"last_update"`
	AverageISPU    float64   `json:"average_ispu"`
}

// StationWithAirQuality combines station and latest air quality data
type StationWithAirQuality struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Type      string    `json:"type"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	ISPU      int       `json:"ispu"`
	Category  string    `json:"category"`
	Color     string    `json:"color"`
	PM25      *float64  `json:"pm25"`
	PM10      *float64  `json:"pm10"`
	CO        *float64  `json:"co"`
	NO2       *float64  `json:"no2"`
	O3        *float64  `json:"o3"`
	SO2       *float64  `json:"so2"`
	HC        *float64  `json:"hc"`
	Timestamp time.Time `json:"timestamp"`
	LastUpdate time.Time `json:"last_update"`
}

// ProvinceStatistic represents statistics per province
type ProvinceStatistic struct {
	Province       string  `json:"province"`
	StationCount   int64   `json:"station_count"`
	AverageISPU    float64 `json:"average_ispu"`
	WorstCategory  string  `json:"worst_category"`
	BestCategory   string  `json:"best_category"`
}

// APIResponse standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
	Meta    *MetaData   `json:"meta,omitempty"`
}

// APIError represents error response
type APIError struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// MetaData represents response metadata
type MetaData struct {
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}
