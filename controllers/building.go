package controllers

import (
	"bul/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	buildingRepository "bul/repository/building"

	"gopkg.in/mgo.v2/bson"
)

func Getdata(res http.ResponseWriter, req *http.Request) {
	var bid = "0000"
	data := GenerateData(bid)
	json.NewEncoder(res).Encode(data)
}

func GetBuildingID(res http.ResponseWriter, req *http.Request) {
	var building models.BuildingName
	var error models.Error

	json.NewDecoder(req.Body).Decode(&building)
	if building.Bid == "" {
		error.Message = "Please fill ID"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(error)
		return
	} else {
		data, _ := buildingRepository.BuildingID(res, req, building.Bid)
		if data != building.Bid {
			error.Message = "Invalid ID"
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(error)
		} else {

			json.NewEncoder(res).Encode(data)
		}

	}

}

func GenerateData(bid string) bson.M {

	bname := buildingRepository.ConnectBname(bid)
	income := buildingRepository.ConnectEstimateIncome(bid)
	land_cost := buildingRepository.ConnectLandCost(bid)
	cost_of_land := buildingRepository.ConnectCostOfLand(bid)
	utility_bill := buildingRepository.ConnectUtilityBill(bid)
	construc_estimate := buildingRepository.ConnectConstrucEstimate(bid)
	more_cost_expense := buildingRepository.ConnectMoreCostExpense(bid)
	operating_cost := buildingRepository.ConnectOperatingCost(bid)
	interest_paid := buildingRepository.ConnectInterestPaid(bid)
	corperate_income_tax := buildingRepository.ConnectCorperateIncomeTax(bid)

	data := bson.M{
		"_id":                  bname.ID,
		"bid":                  bname.Bid,
		"bname":                bname.BName,
		"estimate_income":      income,
		"land_cost":            land_cost,
		"cost_of_land":         cost_of_land,
		"utility_bill":         utility_bill,
		"construc_estimate":    construc_estimate,
		"more_cost_expense":    more_cost_expense,
		"operating_cost":       operating_cost,
		"interest_paid":        interest_paid,
		"corperate_income_tax": corperate_income_tax,
	}

	return data
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
