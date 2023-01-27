package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	artdto "waysgalerry_be/dto/art"
	resultdto "waysgalerry_be/dto/result"
	"waysgalerry_be/models"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerArt struct {
	ArtRepository repositories.ArtRepository
}

func HandlerArt(ArtRepository repositories.ArtRepository) *handlerArt {
	return &handlerArt{ArtRepository}
}

// get all data by user login
func (h *handlerArt) GetArtByUserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// created by
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// get post
	arts, err := h.ArtRepository.GetArtByUserLogin(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: artdto.ArtsResponse{Arts: arts}}
	json.NewEncoder(w).Encode(response)
}

// get all data by user id
func (h *handlerArt) GetArtByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// get post
	arts, err := h.ArtRepository.GetArtByUserId(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: artdto.ArtsResponse{Arts: arts}}
	json.NewEncoder(w).Encode(response)
}

// Create data
func (h *handlerArt) CreateArt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// created by
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	// image
	dataContexErr := r.Context().Value("Error")
	if dataContexErr != true {
		dataContex := r.Context().Value("dataFile")
		filename := dataContex.([]middleware.ImageResult)
		fmt.Println(filename)
		for _, value := range filename {
			// setup data
			art := models.Art{
				ID:        value.PublicID,
				Image:     value.SecureURL,
				CreatedBy: userId,
			}
			// store data
			_, err := h.ArtRepository.CreateArt(art)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
				json.NewEncoder(w).Encode(response)
				return
			}
		}
	}

	// get data
	dataInserted, err := h.ArtRepository.GetArtByUserId(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// fmt.Println(dataInserted)

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: artdto.ArtsResponse{Arts: dataInserted}}
	json.NewEncoder(w).Encode(response)
}
