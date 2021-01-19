package utils

import (
	"jwt/models"
	"os"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
	"github.com/twinj/uuid"
)

func init() {
	gotenv.Load()
}

func GenerateAccessToken(user models.Users) (string, int64, string, error) {
	var err error
	secret := os.Getenv("ACCESS_TOKEN")
	AcExpires := time.Now().Add(time.Minute * 1).Unix()
	AcUid := uuid.NewV4().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"password": user.Password,
		"exp":      AcExpires,
	})

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, AcExpires, AcUid, err

}

func GenerateRefreshToken(user models.Users) (string, int64, string, error) {
	var err error
	secret := os.Getenv("REFRESH_TOKEN")
	RtExpires := time.Now().Add(time.Minute * 30).Unix()
	RtUid := uuid.NewV4().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"email":    user.Email,
		"password": user.Password,
		"exp":      RtExpires,
	})

	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, RtExpires, RtUid, err

}

func VertifyEmail(token string, secret string) (string, string) {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	var email, pass string
	for key, val := range claims {
		if key == "email" {
			email = val.(string)
		}
		if key == "password" {
			pass = val.(string)
		}
	}
	return email, pass
}
