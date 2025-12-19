package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ispu-monitoring/backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	var dsn string
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		dsn = dbURL
	} else {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSLMODE"),
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Only run migrations if explicitly enabled (to avoid slow startup on serverless)
	if os.Getenv("RUN_MIGRATIONS") == "true" {
		log.Println("Running database migrations...")
		err = db.AutoMigrate(
			&model.Station{},
			&model.AirQuality{},
			&model.ISPUCategory{},
		)

		if err != nil {
			return nil, fmt.Errorf("failed to migrate database: %w", err)
		}

		log.Println("Database migrated successfully")

		// Seed categories if empty
		var count int64
		db.Model(&model.ISPUCategory{}).Count(&count)
		if count == 0 {
			SeedCategories(db)
		}
	}

	log.Println("Database connected successfully")

	return db, nil
}

func SeedCategories(db *gorm.DB) {
	categories := []model.ISPUCategory{
		{MinValue: 0, MaxValue: ptrInt(50), Category: "Baik", Description: "Kualitas udara yang sangat baik", Color: "#00e400"},
		{MinValue: 51, MaxValue: ptrInt(100), Category: "Sedang", Description: "Kualitas udara dapat diterima", Color: "#ffff00"},
		{MinValue: 101, MaxValue: ptrInt(200), Category: "Tidak Sehat", Description: "Tidak sehat untuk kelompok sensitif", Color: "#ff7e00"},
		{MinValue: 201, MaxValue: ptrInt(300), Category: "Sangat Tidak Sehat", Description: "Memberikan efek lebih berat", Color: "#ff0000"},
		{MinValue: 301, MaxValue: nil, Category: "Berbahaya", Description: "Berbahaya untuk semua", Color: "#8f3f97"},
	}

	for _, cat := range categories {
		db.Create(&cat)
	}

	log.Println("Categories seeded successfully")
}

func ptrInt(i int) *int {
	return &i
}
