package router

import (
	"buildapi2/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteaMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies",controller.DeleteAllMovies).Methods("DELETE")

	return router
}