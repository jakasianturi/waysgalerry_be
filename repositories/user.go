package repositories

import (
	"waysgalerry_be/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// user repository interface
// interaction to database
type UserRepository interface {
	GetUser(ID int) (models.User, error)
	FindArtsByUserId(ID int) ([]models.Art, error)
	FindPostsByUserId(ID int) ([]models.PostUserResponse, error)
	GetUserDetailByLogin(ID int) (models.UserResponse, error)
	GetUserDetailById(ID int) (models.UserResponse, error)
	UpdateProfile(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

// Get data by ID
func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}

// get data art by user id
func (r *repository) FindArtsByUserId(ID int) ([]models.Art, error) {
	var art []models.Art
	err := r.db.Where("created_by=?", ID).Find(&art).Error

	return art, err
}

// Find post by ID User
func (r *repository) FindPostsByUserId(ID int) ([]models.PostUserResponse, error) {
	var post []models.PostUserResponse
	err := r.db.Debug().Where("created_by=?", ID).Preload(clause.Associations).Find(&post).Error

	return post, err
}

// Get data by ID
func (r *repository) GetUserDetailByLogin(ID int) (models.UserResponse, error) {
	var user models.UserResponse
	err := r.db.First(&user, ID).Error
	return user, err
}

// Get data by ID
func (r *repository) GetUserDetailById(ID int) (models.UserResponse, error) {
	var user models.UserResponse
	err := r.db.First(&user, ID).Error
	return user, err
}

// Update data
func (r *repository) UpdateProfile(user models.User) (models.User, error) {
	err := r.db.Model(&user).Updates(user).Error
	return user, err
}
