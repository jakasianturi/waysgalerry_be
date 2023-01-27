package handlers

import (
	"encoding/json"
	"net/http"
	followdto "waysgalerry_be/dto/follow"
	resultdto "waysgalerry_be/dto/result"
	"waysgalerry_be/models"
	"waysgalerry_be/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerFollow struct {
	FollowRepository repositories.FollowRepository
}

func HandlerFollow(FollowRepository repositories.FollowRepository) *handlerFollow {
	return &handlerFollow{FollowRepository}
}

// Create data
func (h *handlerFollow) Follow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request sesuai dengan register request
	request := new(followdto.Follow)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// validation dari request
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// following
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	// follower
	// setup data
	follow := models.Follow{
		Follower:  userId,
		Following: request.Following,
	}

	// store data
	data, err := h.FollowRepository.Follow(follow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: followdto.FollowResponse{Follow: data}}
	json.NewEncoder(w).Encode(response)
}

// get data by ID
func (h *handlerFollow) UnFollow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request sesuai dengan register request
	request := new(followdto.Follow)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// validation dari request
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// follower
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	// get data
	delete, err := h.FollowRepository.UnFollow(userId, request.Following)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: followdto.FollowResponse{Follow: delete}}
	json.NewEncoder(w).Encode(response)
}
