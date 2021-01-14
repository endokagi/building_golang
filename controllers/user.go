package controllers

import (
	"encoding/json"
	"fmt"
	"jwt/models"
	userRepository "jwt/repository/user"
	"jwt/utils"
	"net/http"
	"os"
	"strings"

	// "time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
}

func Login(response http.ResponseWriter, request *http.Request) {
	var user models.Users
	var error models.Error
	var tokenDetail models.Tokens
	// var token driver.Token
	json.NewDecoder(request.Body).Decode(&user)
	if user.Email == "" || user.Password == "" {
		error.Message = "Invalid Account"
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(error)
		return
	} else {

		user, error = userRepository.FindUser(response, request, user)
		if error.Message == "" {
			Token, AcExpire, AccessUuid, _ := utils.GenerateAccessToken(user)
			refreshToken, RtExpire, RefreshUuid, _ := utils.GenerateRefreshToken(user)
			user.AccessToken = Token
			user.RefreshToken = refreshToken

			tokenDetail.AtExpires = AcExpire
			tokenDetail.AccessUuid = AccessUuid

			tokenDetail.RtExpires = RtExpire
			tokenDetail.RefreshUuid = RefreshUuid

			user.Password = ""
			json.NewEncoder(response).Encode(user)

			userRepository.CreateAuth(user.ID, tokenDetail)

		}
	}

}

func Logout(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var errorObject models.Error
		var tokens RefreshToken
		json.NewDecoder(request.Body).Decode(&tokens)

		var RefreshToken string = tokens.RefreshToken

		if len(RefreshToken) != 0 {
			fmt.Println(len(RefreshToken))
			token, error := jwt.Parse(RefreshToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an Error")
				}

				return []byte(os.Getenv("REFRESH_TOKEN")), nil
			})
			if token.Valid {
				next.ServeHTTP(response, request)
			} else {
				response.WriteHeader(http.StatusInternalServerError)
				errorObject.Message = error.Error()
				json.NewEncoder(response).Encode(errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token"
			json.NewEncoder(response).Encode(errorObject)
			return
		}
	})
}

func TokenVerify(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var errorObject models.Error
		authHeader := request.Header.Get("Authorization")

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an Error")
				}
				return []byte(os.Getenv("ACCESS_TOKEN")), nil
			})

			valid, _ := error.(jwt.ValidationError)
			if valid.Errors == jwt.ValidationErrorExpired {
			}

			if token.Valid {
				next.ServeHTTP(response, request)
			} else {
				response.WriteHeader(http.StatusInternalServerError)
				errorObject.Message = error.Error()
				json.NewEncoder(response).Encode(errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token"
			json.NewEncoder(response).Encode(errorObject)
			return
		}
	})
}

var tokens RefreshToken

func ValidRefresh(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var errorObject models.Error

		json.NewDecoder(request.Body).Decode(&tokens)
		var RefreshToken string = tokens.RefreshToken

		if len(RefreshToken) != 0 {
			token, error := jwt.Parse(RefreshToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an Error")
				}

				return []byte(os.Getenv("REFRESH_TOKEN")), nil
			})

			if token.Valid {
				next.ServeHTTP(response, request)
			} else {
				response.WriteHeader(http.StatusInternalServerError)
				errorObject.Message = error.Error()
				json.NewEncoder(response).Encode(errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid token"
			json.NewEncoder(response).Encode(errorObject)
			return
		}
	})
}
