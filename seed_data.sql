-- Seed Data for ISPU Monitoring System
-- This file contains sample station data for testing

-- Note: Make sure to run this after the application has created the tables

-- Sample Stations (20 stations from major cities in Indonesia)
INSERT INTO stations (name, code, type, latitude, longitude, province, city, address, is_active, created_at, updated_at) VALUES
('DKI 1 Bundaran HI', 'DKI1', 'KLHK', -6.195141, 106.822969, 'DKI Jakarta', 'Jakarta Pusat', 'Bundaran Hotel Indonesia', true, NOW(), NOW()),
('DKI 2 Kelapa Gading', 'DKI2', 'KLHK', -6.158197, 106.908806, 'DKI Jakarta', 'Jakarta Utara', 'Kelapa Gading', true, NOW(), NOW()),
('DKI 3 Jagakarsa', 'DKI3', 'KLHK', -6.341667, 106.825278, 'DKI Jakarta', 'Jakarta Selatan', 'Jagakarsa', true, NOW(), NOW()),
('DKI 4 Lubang Buaya', 'DKI4', 'KLHK', -6.289722, 106.891389, 'DKI Jakarta', 'Jakarta Timur', 'Lubang Buaya', true, NOW(), NOW()),
('DKI 5 Kebon Jeruk', 'DKI5', 'KLHK', -6.186944, 106.768056, 'DKI Jakarta', 'Jakarta Barat', 'Kebon Jeruk', true, NOW(), NOW()),

('Jawa Barat 1 Bandung', 'JABAR1', 'KLHK', -6.914744, 107.609810, 'Jawa Barat', 'Bandung', 'Balai Kota Bandung', true, NOW(), NOW()),
('Jawa Barat 2 Cimahi', 'JABAR2', 'KLHK', -6.872185, 107.542320, 'Jawa Barat', 'Cimahi', 'Pusat Kota Cimahi', true, NOW(), NOW()),
('Jawa Barat 3 Bekasi', 'JABAR3', 'KLHK', -6.238270, 106.975571, 'Jawa Barat', 'Bekasi', 'Bekasi Kota', true, NOW(), NOW()),
('Jawa Barat 4 Bogor', 'JABAR4', 'KLHK', -6.595038, 106.816635, 'Jawa Barat', 'Bogor', 'Kebun Raya Bogor', true, NOW(), NOW()),

('Jawa Tengah 1 Semarang', 'JATENG1', 'KLHK', -6.966667, 110.416664, 'Jawa Tengah', 'Semarang', 'Simpang Lima', true, NOW(), NOW()),
('Jawa Tengah 2 Solo', 'JATENG2', 'KLHK', -7.575489, 110.824326, 'Jawa Tengah', 'Surakarta', 'Balai Kota Solo', true, NOW(), NOW()),

('Jawa Timur 1 Surabaya', 'JATIM1', 'KLHK', -7.257472, 112.752090, 'Jawa Timur', 'Surabaya', 'Balai Kota Surabaya', true, NOW(), NOW()),
('Jawa Timur 2 Malang', 'JATIM2', 'KLHK', -7.966620, 112.632632, 'Jawa Timur', 'Malang', 'Alun-alun Malang', true, NOW(), NOW()),

('Bali 1 Denpasar', 'BALI1', 'KLHK', -8.670458, 115.212631, 'Bali', 'Denpasar', 'Renon', true, NOW(), NOW()),
('Bali 2 Sanur', 'BALI2', 'INTEGRASI', -8.695278, 115.262778, 'Bali', 'Denpasar', 'Pantai Sanur', true, NOW(), NOW()),

('Sumatra Utara 1 Medan', 'SUMUT1', 'KLHK', 3.597031, 98.678513, 'Sumatera Utara', 'Medan', 'Lapangan Merdeka', true, NOW(), NOW()),
('Sumatra Selatan 1 Palembang', 'SUMSEL1', 'KLHK', -2.976074, 104.775410, 'Sumatera Selatan', 'Palembang', 'Ampera', true, NOW(), NOW()),

('Kalimantan Timur 1 Balikpapan', 'KALTIM1', 'KLHK', -1.239344, 116.853721, 'Kalimantan Timur', 'Balikpapan', 'Pusat Kota', true, NOW(), NOW()),
('Sulawesi Selatan 1 Makassar', 'SULSEL1', 'KLHK', -5.147665, 119.432732, 'Sulawesi Selatan', 'Makassar', 'Losari', true, NOW(), NOW()),

('Papua 1 Jayapura', 'PAPUA1', 'KLHK', -2.533333, 140.716667, 'Papua', 'Jayapura', 'Hamadi', true, NOW(), NOW());

-- Sample Air Quality Data (random values for demonstration)
-- Note: Adjust timestamps as needed
INSERT INTO air_qualities (station_id, ispu, pm25, pm10, co, no2, o3, so2, timestamp, created_at) VALUES
(1, 45, 12.5, 25.3, 1.2, 15.5, 35.2, 8.3, NOW() - INTERVAL '1 hour', NOW()),
(2, 52, 15.2, 28.5, 1.5, 18.2, 42.1, 9.5, NOW() - INTERVAL '1 hour', NOW()),
(3, 38, 10.1, 22.3, 0.9, 12.3, 28.5, 6.2, NOW() - INTERVAL '1 hour', NOW()),
(4, 67, 22.5, 35.8, 2.1, 25.5, 55.3, 12.8, NOW() - INTERVAL '1 hour', NOW()),
(5, 55, 18.3, 31.2, 1.8, 20.1, 48.2, 10.5, NOW() - INTERVAL '1 hour', NOW()),
(6, 42, 11.8, 24.5, 1.1, 14.8, 32.5, 7.8, NOW() - INTERVAL '1 hour', NOW()),
(7, 48, 13.5, 27.2, 1.4, 16.8, 38.5, 8.9, NOW() - INTERVAL '1 hour', NOW()),
(8, 58, 19.2, 33.5, 1.9, 22.3, 51.2, 11.2, NOW() - INTERVAL '1 hour', NOW()),
(9, 35, 9.2, 20.5, 0.8, 11.2, 25.8, 5.5, NOW() - INTERVAL '1 hour', NOW()),
(10, 62, 20.5, 36.8, 2.0, 24.2, 53.5, 11.8, NOW() - INTERVAL '1 hour', NOW()),
(11, 49, 14.2, 28.8, 1.5, 17.5, 40.2, 9.2, NOW() - INTERVAL '1 hour', NOW()),
(12, 71, 24.8, 39.5, 2.3, 28.5, 62.5, 14.2, NOW() - INTERVAL '1 hour', NOW()),
(13, 56, 18.8, 32.5, 1.8, 21.2, 49.8, 10.8, NOW() - INTERVAL '1 hour', NOW()),
(14, 33, 8.5, 19.2, 0.7, 10.5, 23.5, 5.2, NOW() - INTERVAL '1 hour', NOW()),
(15, 28, 7.2, 17.5, 0.6, 9.2, 21.2, 4.5, NOW() - INTERVAL '1 hour', NOW()),
(16, 65, 21.5, 37.2, 2.1, 26.5, 57.8, 12.5, NOW() - INTERVAL '1 hour', NOW()),
(17, 78, 26.8, 42.5, 2.5, 32.5, 68.5, 15.8, NOW() - INTERVAL '1 hour', NOW()),
(18, 54, 17.5, 30.8, 1.7, 19.8, 47.5, 10.2, NOW() - INTERVAL '1 hour', NOW()),
(19, 46, 12.8, 26.5, 1.3, 15.8, 36.8, 8.5, NOW() - INTERVAL '1 hour', NOW()),
(20, 31, 8.2, 18.8, 0.7, 10.2, 22.8, 4.8, NOW() - INTERVAL '1 hour', NOW());

-- Add historical data for the last 7 days (for station 1 as example)
INSERT INTO air_qualities (station_id, ispu, pm25, pm10, co, no2, o3, so2, timestamp, created_at)
SELECT 
    1,
    40 + (random() * 40)::int, -- Random ISPU between 40-80
    10 + (random() * 20), -- Random PM2.5
    20 + (random() * 30), -- Random PM10
    0.8 + (random() * 2), -- Random CO
    10 + (random() * 20), -- Random NO2
    20 + (random() * 40), -- Random O3
    5 + (random() * 10), -- Random SO2
    NOW() - (i || ' hours')::interval,
    NOW()
FROM generate_series(1, 168) AS i; -- 168 hours = 7 days

-- Verify the data
SELECT 
    s.name,
    s.code,
    s.province,
    s.city,
    COUNT(aq.id) as reading_count
FROM stations s
LEFT JOIN air_qualities aq ON s.id = aq.station_id
GROUP BY s.id, s.name, s.code, s.province, s.city
ORDER BY s.name;
