package models

import (
	"time"
)

type Hired struct {
	Base
	Title       string       `json:"title" gorm:"type: varchar(255)"`
	Description string       `json:"description" gorm:"type: text"`
	StartDate   time.Time    `json:"startDate" `
	EndDate     time.Time    `json:"endDate"`
	Price       int          `json:"price"`
	Status      string       `json:"status" gorm:"type: varchar(255)"`
	OrderTo     int          `json:"-"`
	OrderBy     int          `json:"-"`
	UserOrderTo UserResponse ` gorm:"foreignkey:OrderTo" json:"orderTo"`
	UserOrderBy UserResponse ` gorm:"foreignkey:OrderBy" json:"orderBy"`
}
