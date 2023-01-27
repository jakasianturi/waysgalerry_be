package repositories

import (
	"waysgalerry_be/models"

	"gorm.io/gorm"
)

// post repository interface
// interaction to database
type PostRepository interface {
	FindPosts() ([]models.Post, error)
	FindPostsByCreator(ID int) ([]models.Post, error)
	GetPost(ID int) (models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

// Get all data (Find)
func (r *repository) FindPosts() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Debug().Preload("Photos").Preload("User").Find(&posts).Error

	return posts, err
}

// Get all data (Find) by creator
func (r *repository) FindPostsByCreator(ID int) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Debug().Preload("Photos").Preload("User").Where("created_by=?", ID).Find(&posts).Error

	return posts, err
}

// Get data by ID
func (r *repository) GetPost(ID int) (models.Post, error) {
	var post models.Post
	err := r.db.Preload("Photos").Preload("User").First(&post, ID).Error

	return post, err
}

// Create data
func (r *repository) CreatePost(post models.Post) (models.Post, error) {
	err := r.db.Create(&post).Error

	return post, err
}
