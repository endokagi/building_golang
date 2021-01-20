package main

import (
	"bul/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Getdata).Methods("GET")
	router.HandleFunc("/w", controllers.GetJSON).Methods("GET")
	router.HandleFunc("/id", controllers.GetBuildingID).Methods("POST")
	router.HandleFunc("/insert", controllers.InsertBuilding).Methods("POST")

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}
