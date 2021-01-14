package main

import (
	"jwt/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/token", controllers.ValidRefresh(controllers.GetNewToken)).Methods("POST")
	router.HandleFunc("/profile", controllers.TokenVerify(controllers.GetProfile)).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout(controllers.GetLogout)).Methods("POST")

	router.HandleFunc("/", controllers.GetJSON).Methods("GET")
	router.HandleFunc("/data", controllers.Getdata).Methods("GET")
	router.HandleFunc("/db", controllers.GetDB).Methods("GET")

	http.ListenAndServe(":1234", router)
}
