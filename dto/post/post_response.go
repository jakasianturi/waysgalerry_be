package postdto

import "waysgalerry_be/models"

type DataResponse struct {
	ID          int                 `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Photos      []models.PostImage  `json:"photos"`
	CreatedBy   models.UserResponse `json:"createdBy"`
}

type PostResponse struct {
	Post interface{} `json:"post"`
}
type PostsResponse struct {
	Posts interface{} `json:"posts"`
}
