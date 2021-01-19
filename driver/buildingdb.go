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

func ConnectBuilding() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("building_db"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("building").Collection(os.Getenv("building_name_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Ping(ctx, readpref.Primary())
	return collection
}

func ConnectSomeBuilding(col string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("building_db"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("building").Collection(col)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Ping(ctx, readpref.Primary())
	return collection
}
