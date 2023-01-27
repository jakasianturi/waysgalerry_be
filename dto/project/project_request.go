package projectdto

type CreateProjectRequest struct {
	Description string `json:"description" label:"Description" validate:"required"`
}
