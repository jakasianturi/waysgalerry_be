package authdto

type RegisterResponse struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}
type RegisterEmailValidResponse struct {
	Email string `json:"email"`
}

type LoginResponse struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

type CheckAuthResponse struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type UserResponse struct {
	User interface{} `json:"user"`
}
