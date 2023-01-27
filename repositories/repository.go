package repositories

import "gorm.io/gorm"

// comunicate to gorm DB
type repository struct {
	db *gorm.DB
}
