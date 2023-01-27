package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	hireddto "waysgalerry_be/dto/hired"
	resultdto "waysgalerry_be/dto/result"
	"waysgalerry_be/models"
	"waysgalerry_be/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gookit/validate"
	"github.com/gorilla/mux"
)

type handlerHired struct {
	HiredRepository repositories.HiredRepository
}

func HandlerHired(HiredRepository repositories.HiredRepository) *handlerHired {
	return &handlerHired{HiredRepository}
}

// Create data
func (h *handlerHired) CreateHired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request sesuai dengan register request
	request := new(hireddto.CreateHiredRequest)
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
	start_date, _ := time.Parse("2006-01-02", request.StartDate)
	end_date, _ := time.Parse("2006-01-02", request.EndDate)
	fmt.Println(request.StartDate, start_date)
	// created by
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))
	// setup data
	hired := models.Hired{
		Title:       request.Title,
		Description: request.Description,
		StartDate:   start_date,
		EndDate:     end_date,
		Price:       request.Price,
		OrderBy:     userId,
		OrderTo:     request.OrderTo,
		Status:      "Waiting Accept",
	}

	// store data
	data, err := h.HiredRepository.CreateHired(hired)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// get data
	dataInserted, err := h.HiredRepository.GetHired(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: hireddto.HiredResponse{Hired: dataInserted}}
	json.NewEncoder(w).Encode(response)
}

// get data by ID
func (h *handlerHired) GetHired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// get data
	hired, err := h.HiredRepository.GetHired(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: hireddto.HiredResponse{Hired: hired}}
	json.NewEncoder(w).Encode(response)
}

// update data
func (h *handlerHired) UpdateHired(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request sesuai dengan register request
	request := new(hireddto.UpdateHiredRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// validation
	v := validate.Struct(request)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: v.Errors.All()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	hired, err := h.HiredRepository.GetHired(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// update data
	if request.Status != "" {
		hired.Status = request.Status
	}
	// store data
	data, err := h.HiredRepository.UpdateHired(hired)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: hireddto.HiredResponse{Hired: data}}
	json.NewEncoder(w).Encode(response)
}

// get data by ID
func (h *handlerHired) FindOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// offer by
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// get data
	offer, err := h.HiredRepository.FindOffer(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: hireddto.HiredResponse{Hired: offer}}
	json.NewEncoder(w).Encode(response)
}

// get data by ID
func (h *handlerHired) FindOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// offer to
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// get data
	order, err := h.HiredRepository.FindOrder(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: hireddto.HiredResponse{Hired: order}}
	json.NewEncoder(w).Encode(response)
}
