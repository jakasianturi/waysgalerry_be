package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	PostRoutes(r)
	ArtRoutes(r)
	UserRoutes(r)
	HiredRoutes(r)
	ProjectRoutes(r)
	FollowRoutes(r)
}
