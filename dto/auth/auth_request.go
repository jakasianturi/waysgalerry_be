package authdto

type RegisterRequest struct {
	Email    string `json:"email" label:"Email" validate:"required|email"`
	Password string `json:"password" label:"Password" validate:"required"`
	FullName string `json:"fullName" label:"Full Name" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" label:"Email" validate:"required"`
	Password string `json:"password" label:"Password" validate:"required"`
}
