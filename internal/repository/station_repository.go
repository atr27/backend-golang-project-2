package repository

import (
	"github.com/ispu-monitoring/backend/internal/model"
	"gorm.io/gorm"
)

type StationRepository struct {
	db *gorm.DB
}

func NewStationRepository(db *gorm.DB) *StationRepository {
	return &StationRepository{db: db}
}

func (r *StationRepository) GetAll() ([]model.Station, error) {
	var stations []model.Station
	result := r.db.Where("is_active = ?", true).Order("name ASC").Find(&stations)
	return stations, result.Error
}

func (r *StationRepository) GetByID(id uint) (*model.Station, error) {
	var station model.Station
	result := r.db.First(&station, id)
	return &station, result.Error
}

func (r *StationRepository) GetByProvince(province string) ([]model.Station, error) {
	var stations []model.Station
	result := r.db.Where("province = ? AND is_active = ?", province, true).Find(&stations)
	return stations, result.Error
}

func (r *StationRepository) Create(station *model.Station) error {
	return r.db.Create(station).Error
}

func (r *StationRepository) Update(id uint, station *model.Station) error {
	return r.db.Model(&model.Station{}).Where("id = ?", id).Updates(station).Error
}

func (r *StationRepository) Delete(id uint) error {
	return r.db.Model(&model.Station{}).Where("id = ?", id).Update("is_active", false).Error
}

func (r *StationRepository) CountAll() (int64, error) {
	var count int64
	result := r.db.Model(&model.Station{}).Count(&count)
	return count, result.Error
}

func (r *StationRepository) CountActive() (int64, error) {
	var count int64
	result := r.db.Model(&model.Station{}).Where("is_active = ?", true).Count(&count)
	return count, result.Error
}

func (r *StationRepository) GetProvinces() ([]string, error) {
	var provinces []string
	result := r.db.Model(&model.Station{}).
		Where("is_active = ? AND province IS NOT NULL AND province != ''", true).
		Distinct("province").
		Pluck("province", &provinces)
	return provinces, result.Error
}
