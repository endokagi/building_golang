package userRepository

import (
	"context"
	"encoding/json"

	"jwt/driver"
	"jwt/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
)

func FindUser(response http.ResponseWriter, request *http.Request, member models.Users) (models.Users, models.Error) {
	var collection = driver.ConnectUser()
	var user models.Users = member
	var error models.Error
	err := collection.FindOne(context.TODO(), bson.M{"email": user.Email, "password": user.Password}).Decode(&user)
	if err != nil {
		error.Message = "Invalid Account"
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(error)

	}
	return user, error
}

func CreateAuth(userid primitive.ObjectID, tokenDetail models.Tokens) {
	var tokenDB = driver.ConnectToken()
	at := time.Unix(tokenDetail.AtExpires, 0)
	rt := time.Unix(tokenDetail.RtExpires, 0)
	now := time.Now()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	tokenDB.InsertOne(ctx, bson.M{
		"_id":       tokenDetail.AccessUuid,
		"AtExpires": at,
		"Time_Sub":  now,
	})

	tokenDB.InsertOne(ctx, bson.M{
		"_id":       tokenDetail.RefreshUuid,
		"RtExpires": rt,
		"Time_Sub":  now,
	})
}
