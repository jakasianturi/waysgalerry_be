package followdto

type Follow struct {
	Following int `json:"following" validate:"required"`
}
