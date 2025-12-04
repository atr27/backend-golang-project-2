package repository

import (
	"time"

	"github.com/ispu-monitoring/backend/internal/model"
	"gorm.io/gorm"
)

type AirQualityRepository struct {
	db *gorm.DB
}

func NewAirQualityRepository(db *gorm.DB) *AirQualityRepository {
	return &AirQualityRepository{db: db}
}

func (r *AirQualityRepository) GetLatestForAllStations() ([]model.AirQuality, error) {
	var results []model.AirQuality
	
	subQuery := r.db.Model(&model.AirQuality{}).
		Select("station_id, MAX(timestamp) as max_timestamp").
		Group("station_id")
	
	err := r.db.
		Joins("INNER JOIN (?) as latest ON air_qualities.station_id = latest.station_id AND air_qualities.timestamp = latest.max_timestamp", subQuery).
		Preload("Station").
		Order("air_qualities.timestamp DESC").
		Find(&results).Error
	
	return results, err
}

func (r *AirQualityRepository) GetLatestByStationID(stationID uint) (*model.AirQuality, error) {
	var airQuality model.AirQuality
	result := r.db.
		Where("station_id = ?", stationID).
		Order("timestamp DESC").
		First(&airQuality)
	return &airQuality, result.Error
}

func (r *AirQualityRepository) GetHistoryByStationID(stationID uint, startDate, endDate time.Time) ([]model.AirQuality, error) {
	var history []model.AirQuality
	result := r.db.
		Where("station_id = ? AND timestamp BETWEEN ? AND ?", stationID, startDate, endDate).
		Order("timestamp DESC").
		Find(&history)
	return history, result.Error
}

func (r *AirQualityRepository) Create(airQuality *model.AirQuality) error {
	return r.db.Create(airQuality).Error
}

func (r *AirQualityRepository) GetAverageISPU() (float64, error) {
	var avg float64
	result := r.db.Model(&model.AirQuality{}).
		Select("COALESCE(AVG(ispu), 0)").
		Scan(&avg)
	return avg, result.Error
}

func (r *AirQualityRepository) GetLatestTimestamp() (time.Time, error) {
	var timestamp time.Time
	result := r.db.Model(&model.AirQuality{}).
		Select("MAX(timestamp)").
		Scan(&timestamp)
	return timestamp, result.Error
}

func (r *AirQualityRepository) GetCategoryDistribution(categories []model.ISPUCategory) (map[string]int, error) {
	distribution := make(map[string]int)
	
	latestData, err := r.GetLatestForAllStations()
	if err != nil {
		return distribution, err
	}
	
	for _, data := range latestData {
		category := GetCategoryForISPU(data.ISPU, categories)
		distribution[category]++
	}
	
	return distribution, nil
}

func GetCategoryForISPU(ispu int, categories []model.ISPUCategory) string {
	for _, cat := range categories {
		if ispu >= cat.MinValue {
			if cat.MaxValue == nil || ispu <= *cat.MaxValue {
				return cat.Category
			}
		}
	}
	return "Unknown"
}
