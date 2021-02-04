package main

import (
	"bul/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/id", controllers.GetBuildingID).Methods("POST")
	router.HandleFunc("/insert", controllers.InsertBuilding).Methods("POST")

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}
