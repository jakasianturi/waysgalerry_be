package routes

import (
	"waysgalerry_be/handlers"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/pkg/mysql"
	"waysgalerry_be/repositories"

	"github.com/gorilla/mux"
)

func ArtRoutes(r *mux.Router) {
	artRepository := repositories.RepositoryArt(mysql.DB)
	h := handlers.HandlerArt(artRepository)

	r.HandleFunc("/arts", middleware.Auth(h.GetArtByUserLogin)).Methods("GET")
	r.HandleFunc("/arts/{id}", middleware.Auth(h.GetArtByUserId)).Methods("GET")
	r.HandleFunc("/upload-arts", middleware.Auth(middleware.UploadMultipleFile(h.CreateArt, "images"))).Methods("POST")
}
