package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"twittergo/middlewares"
	"twittergo/routers"
)

func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
