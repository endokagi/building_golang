package controllers

import (
	"bul/models"
	userRepository "bul/repository/user"
	"encoding/json"
	"net/http"
)

func KubeBname(res http.ResponseWriter, req *http.Request) {
	data := userRepository.KubeBname()
	json.NewEncoder(res).Encode(data)
}

func KubeFindUser(res http.ResponseWriter, req *http.Request) {
	var user models.UserDocker
	var error models.Error

	json.NewDecoder(req.Body).Decode(&user)
	if user.Email == "" || user.Password == "" {
		error.Message = "Please fill email and password"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(error)
		return
	} else {
		users := userRepository.KubeUserID(user)
		if user.Email != users.Email && user.Password != users.Password {
			error.Message = "Invalid Email or Password"
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(error)
		} else {
			json.NewEncoder(res).Encode(users)
		}

	}

}
