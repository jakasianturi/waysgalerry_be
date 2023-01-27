package repositories

import (
	"waysgalerry_be/models"

	"gorm.io/gorm"
)

// auth repository interface
// interaction to database
type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	Getuser(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

// register
func (r *repository) Register(user models.User) (models.User, error) {

	// var err error
	err := r.db.Create(&user).Error

	return user, err
}

// login
func (r *repository) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

// get user
func (r *repository) Getuser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
