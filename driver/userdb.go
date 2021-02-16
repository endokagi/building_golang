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

func ConnectKubeBuilding() *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	auth := options.Credential{
		Username: os.Getenv("KUBE_MONGODB_USERNAME"),
		Password: os.Getenv("KUBE_MONGODB_PASSWORD"),
	}
	clientOptions := options.Client().ApplyURI(os.Getenv("KUBE_MONGODB")).SetAuth(auth)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("KUBE_BUILDING_DATABASE")).Collection(os.Getenv("KUBE_BNAME_COLLECTION"))
	client.Ping(ctx, readpref.Primary())
	return collection
}

func ConnectKubeBuildingUser() *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	auth := options.Credential{
		Username: os.Getenv("KUBE_MONGODB_USERNAME"),
		Password: os.Getenv("KUBE_MONGODB_PASSWORD"),
	}
	clientOptions := options.Client().ApplyURI(os.Getenv("KUBE_MONGODB")).SetAuth(auth)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database(os.Getenv("KUBE_BUILDING_DATABASE")).Collection("users")
	client.Ping(ctx, readpref.Primary())
	return collection
}
