package router

import (
	"github.com/gorilla/mux"
	controller "github.com/harshRishi/mongoapis/Controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Our Routes
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateOneMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/delete-all-movies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
