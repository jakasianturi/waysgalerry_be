package projectdto

import "waysgalerry_be/models"

type DataResponse struct {
	Description string             `json:"description"`
	Photos      []models.PostImage `json:"photos"`
	Hired       models.Hired       `json:"hired"`
}
type ProjectResponse struct {
	Project interface{} `json:"project"`
}
