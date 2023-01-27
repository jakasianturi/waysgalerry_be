package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	resultdto "waysgalerry_be/dto/result"
	userdto "waysgalerry_be/dto/user"
	"waysgalerry_be/models"
	"waysgalerry_be/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

// get data by ID
func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// created by
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// get data
	user, err := h.UserRepository.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: userdto.UserResponse{User: convertUserResponse(user)}}
	json.NewEncoder(w).Encode(response)
}

// get data by ID
func (h *handlerUser) GetUserDetailByLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// created by
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// get data profile
	user, err := h.UserRepository.GetUserDetailByLogin(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// get data profile
	arts, err := h.UserRepository.FindArtsByUserId(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// get data profile
	posts, err := h.UserRepository.FindPostsByUserId(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: userdto.UserDetailResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Greeting: user.Greeting,
		Avatar:   user.Avatar,
		Posts:    posts,
		Arts:     arts,
	}}
	json.NewEncoder(w).Encode(response)
}

// get data by ID
func (h *handlerUser) GetUserDetailById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// get data profile
	user, err := h.UserRepository.GetUserDetailById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// get data profile
	arts, err := h.UserRepository.FindArtsByUserId(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// get data profile
	posts, err := h.UserRepository.FindPostsByUserId(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: userdto.UserDetailResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Greeting: user.Greeting,
		Avatar:   user.Avatar,
		Posts:    posts,
		Arts:     arts,
	}}
	json.NewEncoder(w).Encode(response)
}

// update data
func (h *handlerUser) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request
	request := userdto.UpdateUserRequest{
		Greeting: r.FormValue("greeting"),
		FullName: r.FormValue("fullName"),
	}

	// fmt.Println(request)
	// validation
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// created by
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// update data
	if request.FullName != "" {
		user.FullName = request.FullName
	}
	if request.Greeting != "" {
		user.Greeting = request.Greeting
	}
	// image
	dataContexErr := r.Context().Value("Error")
	if dataContexErr != true {
		dataContex := r.Context().Value("dataFile")
		avatar := dataContex.(string)
		user.Avatar = avatar
	}

	// store data
	data, err := h.UserRepository.UpdateProfile(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: userdto.UserResponse{User: convertUserResponse(data)}}
	json.NewEncoder(w).Encode(response)
}
func convertUserResponse(r models.User) userdto.UpdateResponse {
	return userdto.UpdateResponse{
		Avatar:   r.Avatar,
		Greeting: r.Greeting,
		FullName: r.FullName,
	}
}
