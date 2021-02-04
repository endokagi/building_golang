package driver

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	gotenv.Load()
}

func ConnectUser() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("USER_DB"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("sample_mflix").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Ping(ctx, readpref.Primary())
	return collection
}
