package hireddto

type CreateHiredRequest struct {
	Title       string `json:"title" label:"Title" validate:"required"`
	Description string `json:"description" label:"Description" validate:"required"`
	StartDate   string `json:"startDate" label:"Start Project" validate:"required"`
	EndDate     string `json:"endDate" label:"End Project" validate:"required"`
	Price       int    `json:"price" label:"Price" validate:"required"`
	OrderTo     int    `json:"orderTo" label:"OrderTo" validate:"required"`
}
type UpdateHiredRequest struct {
	Status string `json:"status" label:"Status" validate:"required"`
}
