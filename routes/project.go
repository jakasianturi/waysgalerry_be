package routes

import (
	"waysgalerry_be/handlers"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/pkg/mysql"
	"waysgalerry_be/repositories"

	"github.com/gorilla/mux"
)

func ProjectRoutes(r *mux.Router) {
	projectRepository := repositories.RepositoryProject(mysql.DB)
	h := handlers.HandlerProject(projectRepository)

	r.HandleFunc("/project/{hired_id}", middleware.Auth(middleware.UploadMultipleFile(h.CreateProject, "photos"))).Methods("POST")
	r.HandleFunc("/project/detail/{hired_id}", middleware.Auth(h.GetProject)).Methods("GET")
}
