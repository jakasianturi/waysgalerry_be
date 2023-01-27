package postdto

type CreatePostRequest struct {
	Title       string `json:"title" label:"Title" validate:"required"`
	Description string `json:"description" label:"Description" validate:"required"`
}
