package repositories

import (
	"waysgalerry_be/models"

	"gorm.io/gorm"
)

// art repository interface
// interaction to database
type ArtRepository interface {
	GetArtByUserLogin(ID int) ([]models.Art, error)
	GetArtByUserId(ID int) ([]models.Art, error)
	CreateArt(art models.Art) (models.Art, error)
}

func RepositoryArt(db *gorm.DB) *repository {
	return &repository{db}
}

// get data
func (r *repository) GetArtByUserLogin(ID int) ([]models.Art, error) {
	var art []models.Art
	err := r.db.Debug().Where("created_by=?", ID).Find(&art).Error

	return art, err
}

// get data
func (r *repository) GetArtByUserId(ID int) ([]models.Art, error) {
	var art []models.Art
	err := r.db.Where("created_by=?", ID).Find(&art).Error

	return art, err
}

// Create data
func (r *repository) CreateArt(art models.Art) (models.Art, error) {
	err := r.db.Create(&art).Error

	return art, err
}
