package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ispu-monitoring/backend/internal/config"
	"github.com/ispu-monitoring/backend/internal/handler"
	"github.com/ispu-monitoring/backend/internal/middleware"
	"github.com/ispu-monitoring/backend/internal/repository"
	"github.com/ispu-monitoring/backend/internal/service"
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

	// Initialize Redis
	redisClient := config.InitRedis()

	// Initialize repositories
	stationRepo := repository.NewStationRepository(db)
	airQualityRepo := repository.NewAirQualityRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	// Initialize services
	stationService := service.NewStationService(stationRepo, redisClient)
	airQualityService := service.NewAirQualityService(airQualityRepo, stationRepo, redisClient)
	dashboardService := service.NewDashboardService(stationRepo, airQualityRepo, categoryRepo, redisClient)

	// Initialize handlers
	stationHandler := handler.NewStationHandler(stationService, dashboardService)
	airQualityHandler := handler.NewAirQualityHandler(airQualityService)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	// Initialize Gin router
	r := gin.Default()

	// CORS configuration
	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Middleware
	r.Use(middleware.RequestLogger())
	r.Use(middleware.ErrorHandler())

	// API Routes
	api := r.Group("/api/v1")
	{
		// Health check
		api.GET("/health", handler.HealthCheck)

		// Station endpoints
		stations := api.Group("/stations")
		{
			stations.GET("", stationHandler.GetAllStations)
			stations.GET("/:id", stationHandler.GetStationByID)
			stations.GET("/:id/latest", stationHandler.GetStationLatestData)
			stations.POST("", stationHandler.CreateStation)
			stations.PUT("/:id", stationHandler.UpdateStation)
			stations.DELETE("/:id", stationHandler.DeleteStation)
		}

		// Air quality endpoints
		airQuality := api.Group("/air-quality")
		{
			airQuality.GET("/latest", airQualityHandler.GetLatestData)
			airQuality.GET("/station/:id", airQualityHandler.GetStationHistory)
			airQuality.POST("", airQualityHandler.InsertAirQuality)
		}

		// Dashboard endpoints
		dashboard := api.Group("/dashboard")
		{
			dashboard.GET("/overview", dashboardHandler.GetOverview)
			dashboard.GET("/statistics", dashboardHandler.GetStatistics)
		}

		// Map endpoints
		maps := api.Group("/map")
		{
			maps.GET("/stations", stationHandler.GetMapStations)
		}

		// Categories
		api.GET("/categories", dashboardHandler.GetCategories)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
