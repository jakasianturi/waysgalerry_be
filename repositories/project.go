package repositories

import (
	"waysgalerry_be/models"

	"gorm.io/gorm"
)

// post repository interface
// interaction to database
type ProjectRepository interface {
	GetProject(ID int) (models.Project, error)
	CreateProject(Project models.Project) (models.Project, error)
}

func RepositoryProject(db *gorm.DB) *repository {
	return &repository{db}
}

// Get data by ID
func (r *repository) GetProject(ID int) (models.Project, error) {
	var project models.Project
	err := r.db.Preload("Photos").Preload("Hired.UserOrderBy").Preload("Hired.UserOrderTo").Preload("Hired").Where("hired_id=?", ID).First(&project).Error

	return project, err
}

// Create data
func (r *repository) CreateProject(project models.Project) (models.Project, error) {
	err := r.db.Create(&project).Error

	return project, err
}
