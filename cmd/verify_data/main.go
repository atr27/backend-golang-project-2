package main

import (
	"fmt"
	"log"

	"github.com/ispu-monitoring/backend/internal/config"
	"github.com/ispu-monitoring/backend/internal/model"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Count stations
	var stationCount int64
	db.Model(&model.Station{}).Count(&stationCount)
	fmt.Printf("Total Stations: %d\n", stationCount)

	// Count air quality records
	var airQualityCount int64
	db.Model(&model.AirQuality{}).Count(&airQualityCount)
	fmt.Printf("Total Air Quality Records: %d\n", airQualityCount)

	// Count ISPU categories
	var categoryCount int64
	db.Model(&model.ISPUCategory{}).Count(&categoryCount)
	fmt.Printf("Total ISPU Categories: %d\n", categoryCount)

	// Show sample stations
	var stations []model.Station
	db.Limit(5).Find(&stations)
	fmt.Println("\nSample Stations:")
	for _, station := range stations {
		fmt.Printf("  - %s (%s) - %s, %s\n", station.Name, station.Code, station.City, station.Province)
	}

	// Show air quality stats by station
	type StationStats struct {
		StationName string
		StationCode string
		RecordCount int64
	}
	var stats []StationStats
	db.Raw(`
		SELECT s.name as station_name, s.code as station_code, COUNT(aq.id) as record_count
		FROM stations s
		LEFT JOIN air_qualities aq ON s.id = aq.station_id
		GROUP BY s.id, s.name, s.code
		ORDER BY s.name
	`).Scan(&stats)

	fmt.Println("\nAir Quality Records per Station:")
	for _, stat := range stats {
		fmt.Printf("  - %s (%s): %d records\n", stat.StationName, stat.StationCode, stat.RecordCount)
	}
}
