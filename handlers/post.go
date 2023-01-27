package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	postdto "waysgalerry_be/dto/post"
	resultdto "waysgalerry_be/dto/result"
	"waysgalerry_be/models"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gookit/validate"
	"github.com/gorilla/mux"
)

type handlerPost struct {
	PostRepository repositories.PostRepository
}

func HandlerPost(PostRepository repositories.PostRepository) *handlerPost {
	return &handlerPost{PostRepository}
}

// get all data
func (h *handlerPost) FindPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get post
	posts, err := h.PostRepository.FindPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: postdto.PostsResponse{Posts: posts}}
	json.NewEncoder(w).Encode(response)
}

// get data by ID
func (h *handlerPost) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// get data
	post, err := h.PostRepository.GetPost(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: postdto.PostResponse{Post: post}}
	json.NewEncoder(w).Encode(response)
}

// Create data
func (h *handlerPost) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request
	request := postdto.CreatePostRequest{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	// validation
	v := validate.Struct(request)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: v.Errors.All()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// created by
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	// image
	dataContexErr := r.Context().Value("Error")
	var photos []models.PostImage
	if dataContexErr != true {
		dataContex := r.Context().Value("dataFile")
		filename := dataContex.([]middleware.ImageResult)
		photos = make([]models.PostImage, len(filename))
		for i, value := range filename {
			photos[i] = models.PostImage{ID: value.PublicID, Image: value.SecureURL}
		}
	}
	// setup data
	post := models.Post{
		Title:       request.Title,
		Description: request.Description,
		CreatedBy:   userId,
		Photos:      photos,
	}

	// store data
	data, err := h.PostRepository.CreatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// get data
	dataInserted, err := h.PostRepository.GetPost(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: postdto.PostResponse{Post: dataInserted}}
	json.NewEncoder(w).Encode(response)
}
