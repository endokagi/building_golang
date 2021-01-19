package buildingRepository

import (
	"context"
	"jwt/driver"
	"jwt/models"
	"net/http"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	gotenv.Load()
}

func InsertBuildingName(res http.ResponseWriter, req *http.Request, bname models.BuildingName) {
	var collection = driver.ConnectBuilding()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// interest_paid_collection = driver.ConnectColBuilding(os.Getenv("interest_paid_collection"))
	// corporate_income_tax_collection = driver.ConnectColBuilding(os.Getenv("corporate_income_tax_collection"))

	collection.InsertOne(ctx, bson.M{
		"_id":   bname.ID,
		"bid":   bname.Bid,
		"bname": bname.BName,
	})
}

func InsertEstimateIncome(res http.ResponseWriter, req *http.Request, income models.EstimateIncome, data []interface{}) {
	var collection = driver.ConnectColBuilding(os.Getenv("estimate_income_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bson.M{
		"_id":  income.ID,
		"bid":  income.Bid,
		"area": data,
	})
}

func InsertLandCost(res http.ResponseWriter, req *http.Request, land_cost models.LandCost) {
	var collection = driver.ConnectColBuilding(os.Getenv("cost_of_land_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bson.M{
		"_id":             land_cost.ID,
		"bid":             land_cost.Bid,
		"title":           land_cost.Title,
		"nland":           land_cost.N_land,
		"nPerLand":        land_cost.N_Perland,
		"tranferLandCost": land_cost.TransterLandCost,
	})
}

func InsertCostOfLand(res http.ResponseWriter, req *http.Request, cost_of_land models.CostOfLand) {
	var collection = driver.ConnectColBuilding(os.Getenv("estimate_income_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bson.M{
		"_id":        cost_of_land.ID,
		"bid":        cost_of_land.Bid,
		"title":      cost_of_land.Title,
		"costOfLand": cost_of_land.CostOfLand,
	})
}

func InsertUtilityBill(res http.ResponseWriter, req *http.Request, utility_bill models.UtilityBill, data []interface{}) {
	var collection = driver.ConnectColBuilding(os.Getenv("utility_bill_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bson.M{
		"_id":         utility_bill.ID,
		"bid":         utility_bill.Bid,
		"title":       utility_bill.Title,
		"utilityBill": data,
	})
}

func InsertConstrucEstimate(res http.ResponseWriter, req *http.Request, construc_estimate models.ConstrucEstimate) {
	var collection = driver.ConnectColBuilding(os.Getenv("construction_cost_estimate_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bson.M{
		"_id":              construc_estimate.ID,
		"bid":              construc_estimate.Bid,
		"title":            construc_estimate.Title,
		"nBuildingArea":    construc_estimate.N_BuildingArea,
		"nPerBuildingArea": construc_estimate.N_PerBuildingArea,
	})
}

func InsertMoreCostExpense(res http.ResponseWriter, req *http.Request, more_cost_expense models.MoreCostExpense, data []interface{}) {
	var collection = driver.ConnectColBuilding(os.Getenv("more_cost_expense_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bson.M{
		"_id":                  more_cost_expense.ID,
		"bid":                  more_cost_expense.Bid,
		"title":                more_cost_expense.Title,
		"ownershipTransferFee": more_cost_expense.OwnershipTransferFee,
		"moreCostExpense":      data,
	})
}

func InsertOperatingCost(res http.ResponseWriter, req *http.Request, operating_cost models.OperatingCost, data []interface{}) {
	var collection = driver.ConnectColBuilding(os.Getenv("operating_cost_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bson.M{
		"_id":               operating_cost.ID,
		"bid":               operating_cost.Bid,
		"title":             operating_cost.Title,
		"commission":        operating_cost.Commission,
		"creditUsageFee":    operating_cost.CreditUsageFee,
		"moreOperatingCost": data,
	})
}

func InsertInterestPaid(res http.ResponseWriter, req *http.Request, interest_paid models.InterestPaid) {
	var collection = driver.ConnectColBuilding(os.Getenv("construction_cost_estimate_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bson.M{
		"_id":          interest_paid.ID,
		"bid":          interest_paid.Bid,
		"title":        interest_paid.Title,
		"interestPaid": interest_paid.InterestPaid,
	})
}

func InsertCorporateIncomeTax(res http.ResponseWriter, req *http.Request, corporate_income models.CorporateIncomeTax) {
	var collection = driver.ConnectColBuilding(os.Getenv("construction_cost_estimate_collection"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection.InsertOne(ctx, bson.M{
		"_id":                corporate_income.ID,
		"bid":                corporate_income.Bid,
		"title":              corporate_income.Title,
		"corporateIncomeTax": corporate_income.CorporateIncomeTax,
	})
}
