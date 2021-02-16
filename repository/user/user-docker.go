package userRepository

import (
	"bul/driver"
	"bul/models"
	"context"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func KubeBname() []bson.M {
	var collection = driver.ConnectKubeBuilding()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var users []bson.M
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}
	return users
}

func KubeUserID(user models.UserDocker) models.UserDocker {
	var collection = driver.ConnectKubeBuildingUser()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"email": user.Email, "password": user.Password}).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
