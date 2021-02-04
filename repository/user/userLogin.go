package userRepository

import (
	"bul/driver"
	"bul/models"
	"context"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// UserID is findOne user by email & password
func UserID(user models.Login) models.Login {
	var collection = driver.ConnectUser()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"email": user.Email, "password": user.Password}).Decode(&user)
	if err != nil {
		log.Fatal(err)

	}
	return user
}
