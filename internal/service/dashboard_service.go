package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ispu-monitoring/backend/internal/model"
	"github.com/ispu-monitoring/backend/internal/repository"
	"github.com/redis/go-redis/v9"
)

type DashboardService struct {
	stationRepo  *repository.StationRepository
	airQualityRepo *repository.AirQualityRepository
	categoryRepo *repository.CategoryRepository
	redis        *redis.Client
}

func NewDashboardService(
	stationRepo *repository.StationRepository,
	airQualityRepo *repository.AirQualityRepository,
	categoryRepo *repository.CategoryRepository,
	redis *redis.Client,
) *DashboardService {
	return &DashboardService{
		stationRepo:    stationRepo,
		airQualityRepo: airQualityRepo,
		categoryRepo:   categoryRepo,
		redis:          redis,
	}
}

func (s *DashboardService) GetOverview() (*model.DashboardOverview, error) {
	// Try cache first
	if s.redis != nil {
		cacheKey := "dashboard:overview"
		ctx := context.Background()
		
		cached, err := s.redis.Get(ctx, cacheKey).Result()
		if err == nil {
			var overview model.DashboardOverview
			if err := json.Unmarshal([]byte(cached), &overview); err == nil {
				return &overview, nil
			}
		}
	}
	
	// Fetch data
	totalStations, _ := s.stationRepo.CountAll()
	activeStations, _ := s.stationRepo.CountActive()
	averageISPU, _ := s.airQualityRepo.GetAverageISPU()
	lastUpdate, _ := s.airQualityRepo.GetLatestTimestamp()
	
	// Get categories
	categories, _ := s.categoryRepo.GetAll()
	
	// Get category distribution
	distribution, _ := s.airQualityRepo.GetCategoryDistribution(categories)
	
	// Get recent readings
	latestData, _ := s.airQualityRepo.GetLatestForAllStations()
	recentReadings := make([]model.StationWithAirQuality, 0)
	
	for _, data := range latestData {
		if data.Station != nil {
			category, _ := s.categoryRepo.GetCategoryForISPU(data.ISPU)
			reading := model.StationWithAirQuality{
				ID:         data.Station.ID,
				Name:       data.Station.Name,
				Code:       data.Station.Code,
				Type:       data.Station.Type,
				Latitude:   data.Station.Latitude,
				Longitude:  data.Station.Longitude,
				Province:   data.Station.Province,
				City:       data.Station.City,
				Address:    data.Station.Address,
				ISPU:       data.ISPU,
				PM25:       data.PM25,
				PM10:       data.PM10,
				CO:         data.CO,
				NO2:        data.NO2,
				O3:         data.O3,
				SO2:        data.SO2,
				Timestamp:  data.Timestamp,
				LastUpdate: data.Timestamp,
			}
			if category != nil {
				reading.Category = category.Category
				reading.Color = category.Color
			}
			recentReadings = append(recentReadings, reading)
		}
	}
	
	overview := &model.DashboardOverview{
		Summary: model.DashboardSummary{
			TotalStations:  totalStations,
			ActiveStations: activeStations,
			LastUpdate:     lastUpdate,
			AverageISPU:    averageISPU,
		},
		CategoryDistribution: distribution,
		RecentReadings:       recentReadings,
		ProvinceStats:        []model.ProvinceStatistic{},
	}
	
	// Cache for 3 minutes
	if s.redis != nil {
		ctx := context.Background()
		data, _ := json.Marshal(overview)
		s.redis.Set(ctx, "dashboard:overview", data, 3*time.Minute)
	}
	
	return overview, nil
}

func (s *DashboardService) GetCategories() ([]model.ISPUCategory, error) {
	return s.categoryRepo.GetAll()
}

func (s *DashboardService) GetMapStationsData() ([]model.StationWithAirQuality, error) {
	// Try cache first
	if s.redis != nil {
		cacheKey := "map:stations"
		ctx := context.Background()
		
		cached, err := s.redis.Get(ctx, cacheKey).Result()
		if err == nil {
			var stations []model.StationWithAirQuality
			if err := json.Unmarshal([]byte(cached), &stations); err == nil {
				return stations, nil
			}
		}
	}

	// Get categories for coloring
	categories, _ := s.categoryRepo.GetAll()

	// Get latest data for all stations
	latestData, err := s.airQualityRepo.GetLatestForAllStations()
	if err != nil {
		return nil, err
	}

	mapStations := make([]model.StationWithAirQuality, 0)
	for _, data := range latestData {
		if data.Station != nil {
			category := repository.GetCategoryForISPU(data.ISPU, categories)
			stationData := model.StationWithAirQuality{
				ID:        data.Station.ID,
				Name:      data.Station.Name,
				Code:      data.Station.Code,
				Type:      data.Station.Type,
				Latitude:  data.Station.Latitude,
				Longitude: data.Station.Longitude,
				Province:  data.Station.Province,
				City:      data.Station.City,
				Address:   data.Station.Address,
				ISPU:      data.ISPU,
				PM25:      data.PM25,
				PM10:      data.PM10,
				CO:        data.CO,
				NO2:       data.NO2,
				O3:        data.O3,
				SO2:       data.SO2,
				HC:        data.HC,
				Category:  category,
				Timestamp: data.Timestamp,
			}
			
			// Find color for category
			for _, cat := range categories {
				if cat.Category == category {
					stationData.Color = cat.Color
					break
				}
			}
			
			mapStations = append(mapStations, stationData)
		}
	}

	// Cache for 1 minute
	if s.redis != nil {
		ctx := context.Background()
		data, _ := json.Marshal(mapStations)
		s.redis.Set(ctx, "map:stations", data, 1*time.Minute)
	}

	return mapStations, nil
}
