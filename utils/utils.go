package utils

import (
	"context"
	"jwt/models"
	"os"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func SetLookup(collection *mongo.Collection, from string, field string) []primitive.M {

	var path = "$" + from
	lookupStage := bson.D{{"$lookup", bson.D{{"from", from}, {"localField", field}, {"foreignField", field}, {"as", from}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", path}, {"preserveNullAndEmptyArrays", false}}}}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	showLoadedCursor, err := collection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage})
	if err != nil {
		panic(err)
	}
	var showsLoaded []bson.M
	if err = showLoadedCursor.All(ctx, &showsLoaded); err != nil {
		panic(err)
	}

	return showsLoaded
}
