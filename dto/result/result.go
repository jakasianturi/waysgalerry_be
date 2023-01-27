package resultdto

type SuccessResult struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResult struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}
