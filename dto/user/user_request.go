package userdto

type UpdateUserRequest struct {
	Greeting string `json:"greeting"`
	FullName string `json:"fullName"`
}
