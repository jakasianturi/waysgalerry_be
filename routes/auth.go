package routes

import (
	"waysgalerry_be/handlers"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/pkg/mysql"
	"waysgalerry_be/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
