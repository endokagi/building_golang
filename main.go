package main

import (
	"bul/controllers"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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

	router.HandleFunc("/findall", controllers.FindUserDocker).Methods("GET")
	router.HandleFunc("/findid", controllers.DockerUser).Methods("POST")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins)(router))
}
