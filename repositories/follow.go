package repositories

import (
	"waysgalerry_be/models"

	"gorm.io/gorm"
)

// post repository interface
// interaction to database
type FollowRepository interface {
	Follow(follow models.Follow) (models.Follow, error)
	UnFollow(follower int, following int) (models.Follow, error)
}

func RepositoryFollow(db *gorm.DB) *repository {
	return &repository{db}
}

// Create data
func (r *repository) Follow(follow models.Follow) (models.Follow, error) {
	err := r.db.Create(&follow).Error

	return follow, err
}

// Create data
func (r *repository) UnFollow(follower int, following int) (models.Follow, error) {
	var follow models.Follow
	err := r.db.Where("follower = ? AND following = ?", follower, following).Delete(&follow).Error
	return follow, err
}
