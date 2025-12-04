package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ispu-monitoring/backend/internal/model"
	"github.com/ispu-monitoring/backend/internal/repository"
	"github.com/redis/go-redis/v9"
)

type StationService struct {
	repo  *repository.StationRepository
	redis *redis.Client
}

func NewStationService(repo *repository.StationRepository, redis *redis.Client) *StationService {
	return &StationService{
		repo:  repo,
		redis: redis,
	}
}

func (s *StationService) GetAllStations() ([]model.Station, error) {
	// Try cache first
	if s.redis != nil {
		cacheKey := "stations:all"
		ctx := context.Background()
		
		cached, err := s.redis.Get(ctx, cacheKey).Result()
		if err == nil {
			var stations []model.Station
			if err := json.Unmarshal([]byte(cached), &stations); err == nil {
				return stations, nil
			}
		}
		
		// Fetch from database
		stations, err := s.repo.GetAll()
		if err != nil {
			return nil, err
		}
		
		// Cache for 5 minutes
		data, _ := json.Marshal(stations)
		s.redis.Set(ctx, cacheKey, data, 5*time.Minute)
		
		return stations, nil
	}
	
	return s.repo.GetAll()
}

func (s *StationService) GetStationByID(id uint) (*model.Station, error) {
	return s.repo.GetByID(id)
}

func (s *StationService) GetStationsByProvince(province string) ([]model.Station, error) {
	return s.repo.GetByProvince(province)
}

func (s *StationService) CreateStation(station *model.Station) error {
	// Invalidate cache
	if s.redis != nil {
		ctx := context.Background()
		s.redis.Del(ctx, "stations:all")
	}
	return s.repo.Create(station)
}

func (s *StationService) UpdateStation(id uint, station *model.Station) error {
	// Invalidate cache
	if s.redis != nil {
		ctx := context.Background()
		s.redis.Del(ctx, "stations:all")
		s.redis.Del(ctx, fmt.Sprintf("station:%d", id))
	}
	return s.repo.Update(id, station)
}

func (s *StationService) DeleteStation(id uint) error {
	// Invalidate cache
	if s.redis != nil {
		ctx := context.Background()
		s.redis.Del(ctx, "stations:all")
		s.redis.Del(ctx, fmt.Sprintf("station:%d", id))
	}
	return s.repo.Delete(id)
}
