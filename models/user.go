package models

type User struct {
	Base
	FullName string `json:"fullName" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255);unique"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	Greeting string `json:"greeting" gorm:"type: varchar(255)"`
	Avatar   string `json:"avatar" gorm:"type: varchar(255)"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Greeting string `json:"greeting"`
}

func (UserResponse) TableName() string {
	return "users"
}
