package routes

import (
	"waysgalerry_be/handlers"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/pkg/mysql"
	"waysgalerry_be/repositories"

	"github.com/gorilla/mux"
)

func FollowRoutes(r *mux.Router) {
	followRepository := repositories.RepositoryFollow(mysql.DB)
	h := handlers.HandlerFollow(followRepository)

	r.HandleFunc("/follow", middleware.Auth(h.Follow)).Methods("POST")
	r.HandleFunc("/unfollow", middleware.Auth(h.UnFollow)).Methods("DELETE")
}
