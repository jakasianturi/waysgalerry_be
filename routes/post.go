package routes

import (
	"waysgalerry_be/handlers"
	"waysgalerry_be/pkg/middleware"
	"waysgalerry_be/pkg/mysql"
	"waysgalerry_be/repositories"

	"github.com/gorilla/mux"
)

func PostRoutes(r *mux.Router) {
	postRepository := repositories.RepositoryPost(mysql.DB)
	h := handlers.HandlerPost(postRepository)

	r.HandleFunc("/posts", middleware.Auth(h.FindPosts)).Methods("GET")
	r.HandleFunc("/post/{id}", middleware.Auth(h.GetPost)).Methods("GET")
	r.HandleFunc("/post", middleware.Auth(middleware.UploadMultipleFile(h.CreatePost, "photos"))).Methods("POST")
}
