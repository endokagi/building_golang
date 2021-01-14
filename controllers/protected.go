package controllers

import (
	"encoding/json"
	"jwt/models"
	userRepository "jwt/repository/user"
	"jwt/utils"
	"net/http"
	"os"
	"strings"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func GetProfile(response http.ResponseWriter, request *http.Request) {
	var user models.Users

	authHeader := request.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	tokenString := bearerToken[1]
	user.Email, user.Password = utils.VertifyEmail(tokenString, os.Getenv("ACCESS_TOKEN"))
	profile, _ := userRepository.FindUser(response, request, user)
	json.NewEncoder(response).Encode(profile)
}

func GetLogout(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode("Lugout Success!")
}

func GetNewToken(response http.ResponseWriter, request *http.Request) {
	var user models.Users
	var token models.Tokens

	var RefreshToken string = tokens.RefreshToken
	user.Email, user.Password = utils.VertifyEmail(RefreshToken, os.Getenv("REFRESH_TOKEN"))
	profile, _ := userRepository.FindUser(response, request, user)
	access, _, _, _ := utils.GenerateAccessToken(profile)
	refresh, _, _, _ := utils.GenerateRefreshToken(profile)
	token.AccessToken = access
	token.RefreshToken = refresh

	json.NewEncoder(response).Encode(token)

}
