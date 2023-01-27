package models

type Post struct {
	Base
	Title       string       `json:"title" gorm:"type: varchar(255)"`
	Description string       `json:"description" gorm:"type:text"`
	CreatedBy   int          `json:"-"`
	User        UserResponse `gorm:"foreignKey:CreatedBy" json:"createdBy"`
	Photos      []PostImage  `json:"photos" gorm:"foreignkey:PostID"`
}
type PostUserResponse struct {
	Base
	Title       string       `json:"title" gorm:"type: varchar(255)"`
	Description string       `json:"description" gorm:"type:text"`
	CreatedBy   int          `json:"-"`
	User        UserResponse `gorm:"foreignKey:CreatedBy" json:"-"`
	Photos      []PostImage  `json:"photos" gorm:"foreignkey:PostID"`
}

func (PostUserResponse) TableName() string {
	return "posts"
}
