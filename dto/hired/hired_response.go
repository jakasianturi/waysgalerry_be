package hireddto

import "waysgalerry_be/models"

type CreateHiredResponse struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	StartDate   string              `json:"startDate"`
	EndDate     string              `json:"endDate"`
	Price       int                 `json:"price"`
	OrderBy     models.UserResponse `json:"orderBy"`
	OrderTo     models.UserResponse `json:"orderTo"`
}
type UpdateHiredResponse struct {
	Status string `json:"status"`
}

type HiredResponse struct {
	Hired interface{} `json:"hired"`
}
