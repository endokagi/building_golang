package main

import (
	"net/http"
	"os"

	"bul/controllers"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}
