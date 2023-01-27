package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	authdto "waysgalerry_be/dto/auth"
	resultdto "waysgalerry_be/dto/result"
	"waysgalerry_be/models"
	"waysgalerry_be/pkg/bcrypt"
	jwtToken "waysgalerry_be/pkg/jwt"
	"waysgalerry_be/repositories"

	"github.com/gookit/validate"

	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

// user register
func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request sesuai dengan register request
	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	v := validate.Struct(request)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: v.Errors.All()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check email
	usercheck, err := h.AuthRepository.Login(request.Email)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: authdto.RegisterEmailValidResponse{Email: "Email " + usercheck.Email + " sudah digunakan."}}
		json.NewEncoder(w).Encode(response)
		return
	}

	// hashing password sebelum masuk ke DB
	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// setup user sesuai dengan struct di models user
	user := models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: password,
		Avatar:   "",
	}

	// masukkankan data ke database
	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// response after login
	registerResponse := authdto.RegisterResponse{
		Email:    data.Email,
		FullName: data.FullName,
	}
	// success
	w.Header().Set("Content-Type", "application/json")
	response := resultdto.SuccessResult{Status: "success", Data: authdto.UserResponse{User: registerResponse}}
	json.NewEncoder(w).Encode(response)
}

// user login
func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request sesuai dengan login request
	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// set user sesuai dengan struct models user
	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// Check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: "Wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: "Wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 jam expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}
	// respon setelah login
	loginResponse := authdto.LoginResponse{
		Email:    user.Email,
		FullName: user.FullName,
		Token:    token,
	}

	// success
	w.Header().Set("Content-Type", "application/json")
	response := resultdto.SuccessResult{Status: "success", Data: authdto.UserResponse{User: loginResponse}}
	json.NewEncoder(w).Encode(response)

}

// check auth
func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Check User by Id
	user, err := h.AuthRepository.Getuser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	CheckAuthResponse := authdto.CheckAuthResponse{
		FullName: user.FullName,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}

	w.Header().Set("Content-Type", "application/json")
	response := resultdto.SuccessResult{Status: "success", Data: authdto.UserResponse{User: CheckAuthResponse}}
	json.NewEncoder(w).Encode(response)
}
