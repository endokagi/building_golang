package main

import (
	"jwt/controllers"
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
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/token", controllers.ValidRefresh(controllers.GetNewToken)).Methods("POST")
	router.HandleFunc("/profile", controllers.TokenVerify(controllers.GetProfile)).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout(controllers.GetLogout)).Methods("POST")

	router.HandleFunc("/", controllers.Getdata).Methods("GET")
	router.HandleFunc("/id", controllers.GetBuildingID).Methods("POST")

	port := os.Getenv("port")
	http.ListenAndServe(":"+port, router)
}
