package controllers

import (
	"bul/models"
	buildingRepository "bul/repository/building"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func InsertBuilding(res http.ResponseWriter, req *http.Request) {
	var building models.Building

	data, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(data, &building)
	if err != nil {
		fmt.Println("error:", err)
	}

	bname := bson.M{
		"bid":   building.Bid,
		"bname": building.BName,
	}

	income := bson.M{
		"bid":  building.Bid,
		"area": building.Area,
	}

	land_cost := bson.M{
		"bid":              building.Bid,
		"nLand":            building.N_land,
		"nPerland":         building.N_Perland,
		"transferLandCost": building.TransterLandCost,
		"moreLandCost":     building.MoreLandCost,
	}

	cost_of_land := bson.M{
		"bid":        building.Bid,
		"costOfLand": building.CostOfLand,
	}

	utility_bill := bson.M{
		"bid":         building.Bid,
		"utilityBill": building.UtilityBill,
	}

	construc_estimate := bson.M{
		"bid":              building.Bid,
		"nBuildingArea":    building.N_BuildingArea,
		"nPerBuildingArea": building.N_PerBuildingArea,
	}

	more_cost_expense := bson.M{
		"bid":                  building.Bid,
		"ownershipTransferFee": building.OwnershipTransferFee,
		"moreCostExpense":      building.MoreCostExpense,
	}

	operating_cost := bson.M{
		"bid":               building.Bid,
		"commission":        building.Commission,
		"creditUsageFee":    building.CreditUsageFee,
		"moreOperatingCost": building.MoreOperatingCost,
	}

	interest_paid := bson.M{
		"bid":          building.Bid,
		"interestPaid": building.InterestPaid,
	}

	corporate_income_tax := bson.M{
		"bid":                building.Bid,
		"corporateIncomeTax": building.CorporateIncomeTax,
	}

	buildingRepository.InsertbuildingData(res, req,
		bname, income, land_cost, cost_of_land, utility_bill, construc_estimate,
		more_cost_expense, operating_cost, interest_paid, corporate_income_tax)

	// buildingRepository.InsertBuildingName(bname)
	// buildingRepository.InsertEstimateIncome(income)
	// buildingRepository.InsertCostOfLand(cost_of_land)
	// buildingRepository.InsertUtilityBill(utility_bill)
	// buildingRepository.InsertConstrucEstimate(construc_estimate)
	// buildingRepository.InsertMoreCostExpense(more_cost_expense)
	// buildingRepository.InsertOperatingCost(operating_cost)
	// buildingRepository.InsertInterestPaid(interest_paid)
	// buildingRepository.InsertCorporateIncomeTax(corporate_income_tax)
}
