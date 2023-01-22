package router

import (
	"github.com/gorilla/mux"
	"github.com/prajapatiomkar/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie",controller.CreateMovies).Methods("POST")
	router.HandleFunc("/api/movie/{id}",controller.MarkAswatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}",controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie",controller.DeleteAllMovies).Methods("DELETE")
	
	return router
}
