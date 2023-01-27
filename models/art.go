package models

type Art struct {
	ID        string       `json:"id"`
	Image     string       `json:"image" gorm:"type: varchar(255)"`
	CreatedBy int          `json:"-"`
	User      UserResponse `gorm:"foreignKey:CreatedBy" json:"-"`
}
