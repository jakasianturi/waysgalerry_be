package routes

import (
	"waysgalerry_be/handlers"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/pkg/mysql"
	"waysgalerry_be/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/profile", middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user", middleware.Auth(h.GetUserDetailByLogin)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUserDetailById)).Methods("GET")
	r.HandleFunc("/update-profile", middleware.Auth(middleware.UploadFile(h.UpdateProfile, "avatar"))).Methods("PATCH")
}
