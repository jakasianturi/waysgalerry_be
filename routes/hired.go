package routes

import (
	"waysgalerry_be/handlers"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/pkg/mysql"
	"waysgalerry_be/repositories"

	"github.com/gorilla/mux"
)

func HiredRoutes(r *mux.Router) {
	hiredRepository := repositories.RepositoryHired(mysql.DB)
	h := handlers.HandlerHired(hiredRepository)

	r.HandleFunc("/hired", middleware.Auth(h.CreateHired)).Methods("POST")
	r.HandleFunc("/hired/{id}", middleware.Auth(h.GetHired)).Methods("GET")
	r.HandleFunc("/hired/{id}", middleware.Auth(h.UpdateHired)).Methods("PATCH")
	r.HandleFunc("/order", middleware.Auth(h.FindOrder)).Methods("GET")
	r.HandleFunc("/offer", middleware.Auth(h.FindOffer)).Methods("GET")
}
