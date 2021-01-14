package driver

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	gotenv.Load()
}

func ConnectUser() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("user_db"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("sample_mflix").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Ping(ctx, readpref.Primary())

	client.ListDatabaseNames(ctx, bson.M{})
	return collection
}

func ConnectToken() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(os.Getenv("user_token"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Token").Collection("TokenDetail")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Ping(ctx, readpref.Primary())

	client.ListDatabaseNames(ctx, bson.M{})
	return collection
}

// func Connectbuilding(collection string) *mongo.Collection {
// 	clientOptions := options.Client().ApplyURI(os.Getenv("building_db"))
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	estimate_income_collection := client.Database("building").Collection("estimate_income")                       //1
// 	estimate_expenditure_collection := client.Database("building").Collection("estimate_expenditure")             //2
// 	land_cost_collection := client.Database("building").Collection("land_cost")                                   //2.1
// 	cost_of_land_collection := client.Database("building").Collection("cost_of_land")                             //2.2
// 	utility_bill_collection := client.Database("building").Collection("utility_bill")                             //2.3
// 	construction_cost_estimate_collection := client.Database("building").Collection("construction_cost_estimate") //2.4
// 	more_cost_expense_collection := client.Database("building").Collection("more_cost_expense")                   //2.5
// 	operating_cost_collection := client.Database("building").Collection("operating_cost")                         //2.6
// 	interest_paid_collection := client.Database("building").Collection("interest_paid")                           //2.7
// 	corporate_income_tax_collection := client.Database("building").Collection("corporate_income_tax")             //2.8

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client.Ping(ctx, readpref.Primary())

// 	return estimate_income_collection, estimate_expenditure_collection, land_cost_collection, cost_of_land_collection, utility_bill_collection,
// 		construction_cost_estimate_collection, more_cost_expense_collection, operating_cost_collection, interest_paid_collection,
// 		corporate_income_tax_collection
// }
