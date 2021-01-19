package buildingRepository

import (
	"context"
	"encoding/json"
	"jwt/driver"
	"net/http"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	gotenv.Load()
}

func InsertbuildingData(res http.ResponseWriter, req *http.Request,
	bname, income, land_cost, cost_of_land, utility_bill, construc_estimate,
	more_cost_expense, operating_cost, interest_paid, corporate_income bson.M) {
	json.NewEncoder(res).Encode(bname)
	json.NewEncoder(res).Encode(income)
	json.NewEncoder(res).Encode(land_cost)
	json.NewEncoder(res).Encode(cost_of_land)
	json.NewEncoder(res).Encode(utility_bill)
	json.NewEncoder(res).Encode(construc_estimate)
	json.NewEncoder(res).Encode(more_cost_expense)
	json.NewEncoder(res).Encode(operating_cost)
	json.NewEncoder(res).Encode(interest_paid)
	json.NewEncoder(res).Encode(corporate_income)
}

func InsertBuildingName(bname bson.M) {
	var collection = driver.ConnectBuilding()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bname)
}

func InsertEstimateIncome(income bson.M) {
	var collection = driver.ConnectColBuilding(os.Getenv("estimate_income_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, income)
}

func InsertLandCost(land_cost bson.M) {
	var collection = driver.ConnectColBuilding(os.Getenv("land_cost_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, land_cost)
}

func InsertCostOfLand(cost_of_land bson.M) {
	var collection = driver.ConnectColBuilding(os.Getenv("cost_of_land_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, cost_of_land)
}

func InsertUtilityBill(utility_bill bson.M) {
	var collection = driver.ConnectColBuilding(os.Getenv("utility_bill_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, utility_bill)
}

func InsertConstrucEstimate(construc_estimate bson.M) {
	var collection = driver.ConnectColBuilding(os.Getenv("construction_cost_estimate_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, construc_estimate)
}

func InsertMoreCostExpense(more_cost_expense bson.M) {
	var collection = driver.ConnectColBuilding(os.Getenv("more_cost_expense_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, more_cost_expense)
}

func InsertOperatingCost(operating_cost bson.M) {
	var collection = driver.ConnectColBuilding(os.Getenv("operating_cost_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, operating_cost)
}

func InsertInterestPaid(interest_paid bson.M) {
	var collection = driver.ConnectColBuilding(os.Getenv("interest_paid_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, interest_paid)
}

func InsertCorporateIncomeTax(corporate_income bson.M) {
	var collection = driver.ConnectColBuilding(os.Getenv("corporate_income_tax_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, corporate_income)
}
