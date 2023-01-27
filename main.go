package main

import (
	"waysgalerry_be/database"
	go_toolkit "waysgalerry_be/pkg/go-toolkit"
	mysql "waysgalerry_be/pkg/mysql"

	// postgresql "waysgalerry_be/pkg/postgresql"
	"fmt"
	"net/http"
	"os"
	"waysgalerry_be/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// init godotenv
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	// config go-toolkit validator
	go_toolkit.Config()

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	// init gorilla/mux
	r := mux.NewRouter()

	// add path prefix
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// setup allowed Header, Method, and Origin for CORS
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	// get port from .env
	var port = os.Getenv("PORT")

	// Setup server
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
