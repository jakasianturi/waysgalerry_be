package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	projectdto "waysgalerry_be/dto/project"
	resultdto "waysgalerry_be/dto/result"
	"waysgalerry_be/models"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/repositories"

	"github.com/gookit/validate"
	"github.com/gorilla/mux"
)

type handlerProject struct {
	ProjectRepository repositories.ProjectRepository
}

func HandlerProject(ProjectRepository repositories.ProjectRepository) *handlerProject {
	return &handlerProject{ProjectRepository}
}

// get data by ID
func (h *handlerProject) GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params
	hired_id, _ := strconv.Atoi(mux.Vars(r)["hired_id"])
	fmt.Println(hired_id)
	// get data
	project, err := h.ProjectRepository.GetProject(hired_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: projectdto.ProjectResponse{Project: project}}
	json.NewEncoder(w).Encode(response)
}

// Create data
func (h *handlerProject) CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// request
	request := projectdto.CreateProjectRequest{
		Description: r.FormValue("description"),
	}

	// fmt.Println(request)
	// validation
	// validation
	v := validate.Struct(request)
	if !v.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: v.Errors.All()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// params hired
	hired_id, _ := strconv.Atoi(mux.Vars(r)["hired_id"])
	// image
	dataContexErr := r.Context().Value("Error")
	var photos []models.ProjectImage
	if dataContexErr != true {
		dataContex := r.Context().Value("dataFile")
		filename := dataContex.([]middleware.ImageResult)
		photos = make([]models.ProjectImage, len(filename))
		for i, value := range filename {
			photos[i] = models.ProjectImage{ID: value.PublicID, Image: value.SecureURL}
		}
	}
	// setup data
	project := models.Project{
		Description: request.Description,
		HiredId:     hired_id,
		Photos:      photos,
	}

	// store data
	data, err := h.ProjectRepository.CreateProject(project)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// get data
	dataInserted, err := h.ProjectRepository.GetProject(data.HiredId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := resultdto.ErrorResult{Status: "error", Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := resultdto.SuccessResult{Status: "success", Data: projectdto.ProjectResponse{Project: dataInserted}}
	json.NewEncoder(w).Encode(response)
}
