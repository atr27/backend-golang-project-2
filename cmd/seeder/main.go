package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ispu-monitoring/backend/internal/config"
	"github.com/ispu-monitoring/backend/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
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

	log.Println("Starting data seeding...")

	seedStations(db)
	seedAirQuality(db)

	log.Println("Data seeding completed successfully!")
}

func seedStations(db *gorm.DB) {
	stations := []model.Station{
		{Name: "DKI 1 Bundaran HI", Code: "DKI1", Type: "KLHK", Latitude: -6.195141, Longitude: 106.822969, Province: "DKI Jakarta", City: "Jakarta Pusat", Address: "Bundaran Hotel Indonesia", IsActive: true},
		{Name: "DKI 2 Kelapa Gading", Code: "DKI2", Type: "KLHK", Latitude: -6.158197, Longitude: 106.908806, Province: "DKI Jakarta", City: "Jakarta Utara", Address: "Kelapa Gading", IsActive: true},
		{Name: "DKI 3 Jagakarsa", Code: "DKI3", Type: "KLHK", Latitude: -6.341667, Longitude: 106.825278, Province: "DKI Jakarta", City: "Jakarta Selatan", Address: "Jagakarsa", IsActive: true},
		{Name: "DKI 4 Lubang Buaya", Code: "DKI4", Type: "KLHK", Latitude: -6.289722, Longitude: 106.891389, Province: "DKI Jakarta", City: "Jakarta Timur", Address: "Lubang Buaya", IsActive: true},
		{Name: "DKI 5 Kebon Jeruk", Code: "DKI5", Type: "KLHK", Latitude: -6.186944, Longitude: 106.768056, Province: "DKI Jakarta", City: "Jakarta Barat", Address: "Kebon Jeruk", IsActive: true},

		{Name: "Jawa Barat 1 Bandung", Code: "JABAR1", Type: "KLHK", Latitude: -6.914744, Longitude: 107.609810, Province: "Jawa Barat", City: "Bandung", Address: "Balai Kota Bandung", IsActive: true},
		{Name: "Jawa Barat 2 Cimahi", Code: "JABAR2", Type: "KLHK", Latitude: -6.872185, Longitude: 107.542320, Province: "Jawa Barat", City: "Cimahi", Address: "Pusat Kota Cimahi", IsActive: true},
		{Name: "Jawa Barat 3 Bekasi", Code: "JABAR3", Type: "KLHK", Latitude: -6.238270, Longitude: 106.975571, Province: "Jawa Barat", City: "Bekasi", Address: "Bekasi Kota", IsActive: true},
		{Name: "Jawa Barat 4 Bogor", Code: "JABAR4", Type: "KLHK", Latitude: -6.595038, Longitude: 106.816635, Province: "Jawa Barat", City: "Bogor", Address: "Kebun Raya Bogor", IsActive: true},

		{Name: "Jawa Tengah 1 Semarang", Code: "JATENG1", Type: "KLHK", Latitude: -6.966667, Longitude: 110.416664, Province: "Jawa Tengah", City: "Semarang", Address: "Simpang Lima", IsActive: true},
		{Name: "Jawa Tengah 2 Solo", Code: "JATENG2", Type: "KLHK", Latitude: -7.575489, Longitude: 110.824326, Province: "Jawa Tengah", City: "Surakarta", Address: "Balai Kota Solo", IsActive: true},

		{Name: "Jawa Timur 1 Surabaya", Code: "JATIM1", Type: "KLHK", Latitude: -7.257472, Longitude: 112.752090, Province: "Jawa Timur", City: "Surabaya", Address: "Balai Kota Surabaya", IsActive: true},
		{Name: "Jawa Timur 2 Malang", Code: "JATIM2", Type: "KLHK", Latitude: -7.966620, Longitude: 112.632632, Province: "Jawa Timur", City: "Malang", Address: "Alun-alun Malang", IsActive: true},

		{Name: "Bali 1 Denpasar", Code: "BALI1", Type: "KLHK", Latitude: -8.670458, Longitude: 115.212631, Province: "Bali", City: "Denpasar", Address: "Renon", IsActive: true},
		{Name: "Bali 2 Sanur", Code: "BALI2", Type: "INTEGRASI", Latitude: -8.695278, Longitude: 115.262778, Province: "Bali", City: "Denpasar", Address: "Pantai Sanur", IsActive: true},

		{Name: "Sumatra Utara 1 Medan", Code: "SUMUT1", Type: "KLHK", Latitude: 3.597031, Longitude: 98.678513, Province: "Sumatera Utara", City: "Medan", Address: "Lapangan Merdeka", IsActive: true},
		{Name: "Sumatra Selatan 1 Palembang", Code: "SUMSEL1", Type: "KLHK", Latitude: -2.976074, Longitude: 104.775410, Province: "Sumatera Selatan", City: "Palembang", Address: "Ampera", IsActive: true},

		{Name: "Kalimantan Timur 1 Balikpapan", Code: "KALTIM1", Type: "KLHK", Latitude: -1.239344, Longitude: 116.853721, Province: "Kalimantan Timur", City: "Balikpapan", Address: "Pusat Kota", IsActive: true},
		{Name: "Sulawesi Selatan 1 Makassar", Code: "SULSEL1", Type: "KLHK", Latitude: -5.147665, Longitude: 119.432732, Province: "Sulawesi Selatan", City: "Makassar", Address: "Losari", IsActive: true},

		{Name: "Papua 1 Jayapura", Code: "PAPUA1", Type: "KLHK", Latitude: -2.533333, Longitude: 140.716667, Province: "Papua", City: "Jayapura", Address: "Hamadi", IsActive: true},
		{Name: "Kota Tarakan Pamusian", Code: "KALTARA1", Type: "KLHK", Latitude: 3.3, Longitude: 117.633333, Province: "Kalimantan Utara", City: "Tarakan", Address: "Kantor Dinas Lingkungan Hidup", IsActive: true},
	}

	for _, station := range stations {
		var existingStation model.Station
		if err := db.Where("code = ?", station.Code).First(&existingStation).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&station).Error; err != nil {
					log.Printf("Failed to create station %s: %v", station.Name, err)
				} else {
					log.Printf("Created station: %s", station.Name)
				}
			} else {
				log.Printf("Error checking station %s: %v", station.Name, err)
			}
		} else {
			log.Printf("Station %s already exists", station.Name)
		}
	}
}

func seedAirQuality(db *gorm.DB) {
	var stations []model.Station
	if err := db.Find(&stations).Error; err != nil {
		log.Printf("Failed to fetch stations: %v", err)
		return
	}

	// Define ISPU category ranges for better distribution
	ispuCategories := []struct {
		min int
		max int
	}{
		{0, 50},    // Baik
		{51, 100},  // Sedang
		{101, 200}, // Tidak Sehat
		{201, 300}, // Sangat Tidak Sehat
		{301, 400}, // Berbahaya
	}

	for stationIndex, station := range stations {
		// Vary ISPU category for different stations for better testing
		categoryIndex := stationIndex % len(ispuCategories)
		selectedCategory := ispuCategories[categoryIndex]

		// Generate 24 hours of data
		for i := 0; i < 24; i++ {
			timestamp := time.Now().Add(time.Duration(-i) * time.Hour)
			
			// Check if data already exists for this hour
			var count int64
			db.Model(&model.AirQuality{}).
				Where("station_id = ? AND timestamp BETWEEN ? AND ?", 
					station.ID, 
					timestamp.Truncate(time.Hour), 
					timestamp.Truncate(time.Hour).Add(time.Hour)).
				Count(&count)
				
			if count > 0 {
				continue
			}

			var ispu int
			var pm25, pm10, co, no2, o3, so2, hc float64

			// Generate realistic values based on ISPU category
			// Add some variation within the category
			ispu = selectedCategory.min + rand.Intn(selectedCategory.max-selectedCategory.min+1)
			
			// Vary values with small random fluctuations for hourly changes
			baseVariation := float64(i) * 0.5
			
			// Scale pollutant values based on ISPU value
			ispuFactor := float64(ispu) / 100.0
			
			pm25 = (10.0 + ispuFactor*20.0) + rand.Float64()*5 - 2.5 + baseVariation
			pm10 = (15.0 + ispuFactor*35.0) + rand.Float64()*10 - 5 + baseVariation
			so2 = (2.0 + ispuFactor*15.0) + rand.Float64()*3 - 1.5
			co = (1.0 + ispuFactor*5.0) + rand.Float64()*2 - 1
			o3 = (15.0 + ispuFactor*40.0) + rand.Float64()*10 - 5
			no2 = (8.0 + ispuFactor*25.0) + rand.Float64()*5 - 2.5
			hc = (0.5 + ispuFactor*2.0) + rand.Float64()*0.5 - 0.25

			// Ensure no negative values
			if pm25 < 0 {
				pm25 = 0
			}
			if pm10 < 0 {
				pm10 = 0
			}
			if so2 < 0 {
				so2 = 0
			}
			if co < 0 {
				co = 0
			}
			if o3 < 0 {
				o3 = 0
			}
			if no2 < 0 {
				no2 = 0
			}
			if hc < 0 {
				hc = 0
			}

			aq := model.AirQuality{
				StationID: station.ID,
				ISPU:      ispu,
				PM25:      &pm25,
				PM10:      &pm10,
				CO:        &co,
				NO2:       &no2,
				O3:        &o3,
				SO2:       &so2,
				HC:        &hc,
				Timestamp: timestamp,
			}

			if err := db.Create(&aq).Error; err != nil {
				log.Printf("Failed to create air quality data for station %s: %v", station.Name, err)
			}
		}
		log.Printf("Seeded air quality data for station: %s (ISPU range: %d-%d)", 
			station.Name, selectedCategory.min, selectedCategory.max)
	}
}
