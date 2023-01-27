package repositories

import (
	"waysgalerry_be/models"

	"gorm.io/gorm"
)

// post repository interface
// interaction to database
type HiredRepository interface {
	CreateHired(hired models.Hired) (models.Hired, error)
	GetHired(ID int) (models.Hired, error)
	UpdateHired(hired models.Hired) (models.Hired, error)
	FindOffer(ID int) ([]models.Hired, error)
	FindOrder(ID int) ([]models.Hired, error)
}

func RepositoryHired(db *gorm.DB) *repository {
	return &repository{db}
}

// Create data
func (r *repository) CreateHired(hired models.Hired) (models.Hired, error) {
	err := r.db.Create(&hired).Error

	return hired, err
}

// Create data
func (r *repository) UpdateHired(hired models.Hired) (models.Hired, error) {
	err := r.db.Model(&hired).Updates(hired).Error

	return hired, err
}

// Create data
func (r *repository) GetHired(ID int) (models.Hired, error) {
	var hired models.Hired
	err := r.db.Preload("UserOrderBy").Preload("UserOrderTo").First(&hired, ID).Error

	return hired, err
}

// Get all data (Find)
func (r *repository) FindOffer(ID int) ([]models.Hired, error) {
	var hireds []models.Hired
	err := r.db.Debug().Preload("UserOrderBy").Preload("UserOrderTo").Where("order_by=?", ID).Find(&hireds).Error

	return hireds, err
}

// Get all data (Find)
func (r *repository) FindOrder(ID int) ([]models.Hired, error) {
	var hireds []models.Hired
	err := r.db.Debug().Preload("UserOrderBy").Preload("UserOrderTo").Where("order_to=?", ID).Find(&hireds).Error

	return hireds, err
}
