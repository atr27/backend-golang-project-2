package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ispu-monitoring/backend/internal/model"
	"github.com/ispu-monitoring/backend/internal/repository"
	"github.com/redis/go-redis/v9"
)

type AirQualityService struct {
	repo        *repository.AirQualityRepository
	stationRepo *repository.StationRepository
	redis       *redis.Client
}

func NewAirQualityService(repo *repository.AirQualityRepository, stationRepo *repository.StationRepository, redis *redis.Client) *AirQualityService {
	return &AirQualityService{
		repo:        repo,
		stationRepo: stationRepo,
		redis:       redis,
	}
}

func (s *AirQualityService) GetLatestData() ([]model.AirQuality, error) {
	// Try cache first
	if s.redis != nil {
		cacheKey := "air_quality:latest"
		ctx := context.Background()
		
		cached, err := s.redis.Get(ctx, cacheKey).Result()
		if err == nil {
			var data []model.AirQuality
			if err := json.Unmarshal([]byte(cached), &data); err == nil {
				return data, nil
			}
		}
		
		// Fetch from database
		data, err := s.repo.GetLatestForAllStations()
		if err != nil {
			return nil, err
		}
		
		// Cache for 2 minutes
		jsonData, _ := json.Marshal(data)
		s.redis.Set(ctx, cacheKey, jsonData, 2*time.Minute)
		
		return data, nil
	}
	
	return s.repo.GetLatestForAllStations()
}

func (s *AirQualityService) GetStationLatestData(stationID uint) (*model.AirQuality, error) {
	return s.repo.GetLatestByStationID(stationID)
}

func (s *AirQualityService) GetHistoricalData(stationID uint, startDate, endDate time.Time) ([]model.AirQuality, error) {
	return s.repo.GetHistoryByStationID(stationID, startDate, endDate)
}

func (s *AirQualityService) InsertAirQuality(data *model.AirQuality) error {
	// Invalidate cache
	if s.redis != nil {
		ctx := context.Background()
		s.redis.Del(ctx, "air_quality:latest")
		s.redis.Del(ctx, "dashboard:overview")
	}
	return s.repo.Create(data)
}
