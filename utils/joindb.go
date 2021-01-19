package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllBuilding(collection *mongo.Collection) []primitive.M {

	estimate_income := SetBSON("estimate_income", "bid", "bid", "estimate_income")

	estimate_expenditure := SetBSON("estimate_expenditure", "bid", "bid", "estimate_expenditure")

	land_cost := SetBSON("land_cost", "bid", "bid", "land_cost")

	cost_of_land := SetBSON("cost_of_land", "bid", "bid", "cost_of_land")

	utility_bill := SetBSON("utility_bill", "bid", "bid", "utility_bill")

	construction_cost_estimate := SetBSON("construction_cost_estimate", "bid", "bid", "construction_cost_estimate")

	more_cost_expense := SetBSON("more_cost_expense", "bid", "bid", "more_cost_expense")

	operating_cost := SetBSON("operating_cost", "bid", "bid", "operating_cost")

	interest_paid := SetBSON("interest_paid", "bid", "bid", "interest_paid")

	corporate_income_tax := SetBSON("corporate_income_tax", "bid", "bid", "corporate_income_tax")

	project := bson.D{{"$project", bson.M{
		"_id":             1,
		"bid":             1,
		"bname":           1,
		"estimate_income": "$estimate_income",
		"estimate_expenditure": bson.A{
			"$land_cost",
			"$cost_of_land",
			"$utility_bill",
			"$construction_cost_estimate",
			"$more_cost_expense",
			"$operating_cost",
			"$interest_paid",
			"$corporate_income_tax",
		}}}}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	LoadedCursor, err := collection.Aggregate(ctx, mongo.Pipeline{
		estimate_income,
		estimate_expenditure,
		land_cost,
		cost_of_land,
		utility_bill,
		construction_cost_estimate,
		more_cost_expense,
		operating_cost,
		interest_paid,
		corporate_income_tax,
		project})
	var Loaded []bson.M
	if err = LoadedCursor.All(ctx, &Loaded); err != nil {
		panic(err)
	}

	return Loaded
}

func GetBuildingByID(collection *mongo.Collection, id string) interface{} {

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
		"_id":             1,
		"bid":             1,
		"bname":           1,
		"estimate_income": 1,
		"estimate_expenditure": bson.A{
			"$land_cost",
			"$cost_of_land",
			"$utility_bill",
			"$construction_cost_estimate",
			"$more_cost_expense",
			"$operating_cost",
			"$interest_paid",
			"$corporate_income_tax",
		}}}}

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

	var Loaded []bson.M
	if err = LoadedCursor.All(ctx, &Loaded); err != nil {
		panic(err)
	}

	return Loaded
}

func SetBSON(from string, localField string, foreignField string, as string) bson.D {
	bsonD := bson.D{{"$lookup", bson.D{
		{"from", from},
		{"localField", localField},
		{"foreignField", foreignField},
		{"as", as}}}}

	return bsonD
}
