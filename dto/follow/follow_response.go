package followdto

type DataResponse struct {
	Message string `json:"message"`
}

type FollowResponse struct {
	Follow interface{} `json:"follow"`
}
