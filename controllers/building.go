package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jwt/driver"
	"jwt/models"
	"jwt/utils"
	"net/http"

	buildingRepository "jwt/repository/building"

	"gopkg.in/mgo.v2/bson"
)

func Getdata(res http.ResponseWriter, req *http.Request) {
	collection := driver.ConnectBuilding()
	data := utils.GetBuildingByID(collection, "0000")
	json.NewEncoder(res).Encode(data)
}

func GetBuildingID(res http.ResponseWriter, req *http.Request) {
	var building models.BuildingName
	var error models.Error

	json.NewDecoder(req.Body).Decode(&building)
	if building.Bid == "" {
		fmt.Println("bid: ", building.Bid)
		error.Message = "Invalid Account"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(error)
		return
	} else {
		fmt.Println("input: ", building.Bid)
		collection := driver.ConnectBuilding()
		id, _ := buildingRepository.BuildingID(res, req, building)
		data := utils.GetBuildingByID(collection, id)
		json.NewEncoder(res).Encode(data)
	}
}

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

	buildingRepository.InsertBuildingName(bname)
	buildingRepository.InsertEstimateIncome(income)
	buildingRepository.InsertCostOfLand(cost_of_land)
	buildingRepository.InsertUtilityBill(utility_bill)
	buildingRepository.InsertConstrucEstimate(construc_estimate)
	buildingRepository.InsertMoreCostExpense(more_cost_expense)
	buildingRepository.InsertOperatingCost(operating_cost)
	buildingRepository.InsertInterestPaid(interest_paid)
	buildingRepository.InsertCorporateIncomeTax(corporate_income_tax)
}

func GetJSON(res http.ResponseWriter, req *http.Request) {
	var building models.Building
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &building)
	if err != nil {
		fmt.Println("error:", err)
	}
	json.NewEncoder(res).Encode(building)
}
