package utils

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBuildingByID(collection *mongo.Collection, id string) primitive.M {

	matchStage := bson.D{{"$match", bson.D{{"bid", id}}}}

	estimate_income := SetBSON("estimate_income", "bid", "bid", "estimate_income")

	land_cost := SetBSON("land_cost", "bid", "bid", "land_cost")

	cost_of_land := SetBSON("cost_of_land", "bid", "bid", "cost_of_land")

	utility_bill := SetBSON("utility_bill", "bid", "bid", "utility_bill")

	construction_cost_estimate := SetBSON("construction_cost_estimate", "bid", "bid", "construction_cost_estimate")

	more_cost_expense := SetBSON("more_cost_expense", "bid", "bid", "more_cost_expense")

	operating_cost := SetBSON("operating_cost", "bid", "bid", "operating_cost")

	interest_paid := SetBSON("interest_paid", "bid", "bid", "interest_paid")

	corporate_income_tax := SetBSON("corporate_income_tax", "bid", "bid", "corporate_income_tax")

	project := bson.D{{"$project", bson.M{
		"_id":                        1,
		"bid":                        1,
		"bname":                      1,
		"estimate_income":            1,
		"land_cost":                  1,
		"cost_of_land":               1,
		"utility_bill":               1,
		"construction_cost_estimate": 1,
		"more_cost_expense":          1,
		"operating_cost":             1,
		"interest_paid":              1,
		"corporate_income_tax":       1,
	}}}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	LoadedCursor, err := collection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		estimate_income,
		land_cost,
		cost_of_land,
		utility_bill,
		construction_cost_estimate,
		more_cost_expense,
		operating_cost,
		interest_paid,
		corporate_income_tax,
		project})

	var Loaded bson.M
	for LoadedCursor.Next(ctx) {
		if err = LoadedCursor.Decode(&Loaded); err != nil {
			log.Fatal(err)
		}
	}
	return Loaded
}

func SetBSON(from string, localField string, foreignField string, as string) bson.D {
	bsonD := bson.D{{"$lookup", bson.M{
		"from":         from,
		"localField":   localField,
		"foreignField": foreignField,
		"as":           as}}}

	return bsonD
}
