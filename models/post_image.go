package models

type PostImage struct {
	ID     string `json:"id"`
	PostID int    `json:"-"`
	Image  string `json:"image" gorm:"type: varchar(255)"`
}
