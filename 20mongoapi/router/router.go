package router

import (
	"github.com/abhishekkushwahaa/mongoapi/controller"

	"github.com/gorilla/mux"
)


func Router() *mux.Router {
	// Initialize a new router
	router := mux.NewRouter()

	// Route handles & endpoints
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")

	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")

	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")

	router.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")

	return router
}