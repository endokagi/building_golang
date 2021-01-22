package buildingRepository

import (
	"bul/driver"
	"bul/models"
	"context"
	"log"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	gotenv.Load()
}

func ConnectBname(bid string) models.BuildingName {
	var collection = driver.ConnectBuilding()
	var bname models.BuildingName
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&bname)
	if err != nil {
		log.Fatal(err)
	}
	return bname
}

func ConnectEstimateIncome(bid string) models.EstimateIncome {
	var collection = driver.ConnectColBuilding(os.Getenv("estimate_income_collection"))
	var income models.EstimateIncome
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&income)
	if err != nil {
		log.Fatal(err)
	}
	return income
}

// connect to land cost collection
func ConnectLandCost(bid string) models.LandCost {
	var collection = driver.ConnectColBuilding(os.Getenv("land_cost_collection"))
	var land_cost models.LandCost
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&land_cost)
	if err != nil {
		log.Fatal(err)
	}
	return land_cost
}

func ConnectCostOfLand(bid string) models.CostOfLand {
	var collection = driver.ConnectColBuilding(os.Getenv("cost_of_land_collection"))
	var cost_of_land models.CostOfLand
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&cost_of_land)
	if err != nil {
		log.Fatal(err)
	}
	return cost_of_land
}

func ConnectUtilityBill(bid string) models.UtilityBill {
	var collection = driver.ConnectColBuilding(os.Getenv("utility_bill_collection"))
	var utility_bill models.UtilityBill
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&utility_bill)
	if err != nil {
		log.Fatal(err)
	}
	return utility_bill
}

func ConnectConstrucEstimate(bid string) models.ConstrucEstimate {
	var collection = driver.ConnectColBuilding(os.Getenv("construction_cost_estimate_collection"))
	var construcEstimate models.ConstrucEstimate
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&construcEstimate)
	if err != nil {
		log.Fatal(err)
	}
	return construcEstimate
}

func ConnectMoreCostExpense(bid string) models.MoreCostExpense {
	var collection = driver.ConnectColBuilding(os.Getenv("more_cost_expense_collection"))
	var moreCostExpense models.MoreCostExpense
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&moreCostExpense)
	if err != nil {
		log.Fatal(err)
	}
	return moreCostExpense
}

func ConnectOperatingCost(bid string) models.OperatingCost {
	var collection = driver.ConnectColBuilding(os.Getenv("operating_cost_collection"))
	var operating_cost models.OperatingCost
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&operating_cost)
	if err != nil {
		log.Fatal(err)
	}
	return operating_cost
}

func ConnectInterestPaid(bid string) models.InterestPaid {
	var collection = driver.ConnectColBuilding(os.Getenv("interest_paid_collection"))
	var interest_paid models.InterestPaid
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&interest_paid)
	if err != nil {
		log.Fatal(err)
	}
	return interest_paid
}

func ConnectCorperateIncomeTax(bid string) models.CorporateIncomeTax {
	var collection = driver.ConnectColBuilding(os.Getenv("corporate_income_tax_collection"))
	var corperate_income_tax models.CorporateIncomeTax
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&corperate_income_tax)
	if err != nil {
		log.Fatal(err)
	}
	return corperate_income_tax
}
