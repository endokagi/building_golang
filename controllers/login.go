package controllers

import (
	"bul/models"
	userRepository "bul/repository/user"
	"encoding/json"
	"net/http"
)

// Login is function for test login
func Login(res http.ResponseWriter, req *http.Request) {
	var login models.Login
	var error models.Error

	json.NewDecoder(req.Body).Decode(&login)
	if login.Email == "" || login.Password == "" {
		error.Message = "Please fill email and password"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(error)
		return
	} else {
		user := userRepository.UserID(login)
		if user.Email != login.Email && user.Password != login.Password {
			error.Message = "Invalid Email or Password"
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(error)
		} else {
			json.NewEncoder(res).Encode(user)
		}

	}

}
