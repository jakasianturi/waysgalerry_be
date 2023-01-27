package userdto

import "waysgalerry_be/models"

type UpdateResponse struct {
	Avatar   string `json:"avatar"`
	Greeting string `json:"greeting"`
	FullName string `json:"fullName"`
}

type UserDetailResponse struct {
	ID       int                       `json:"id"`
	FullName string                    `json:"fullName"`
	Email    string                    `json:"email"`
	Avatar   string                    `json:"avatar"`
	Greeting string                    `json:"greeting"`
	Posts    []models.PostUserResponse `json:"posts"`
	Arts     []models.Art              `json:"arts"`
}
type UserPostResponse struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Photos      []models.PostImage `json:"photos"`
}
type UserArtResponse struct {
	ID    string `json:"id"`
	Image string `json:"image"`
}

type UserResponse struct {
	User interface{} `json:"user"`
}
