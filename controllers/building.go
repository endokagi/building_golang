package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jwt/driver"
	"jwt/models"
	"jwt/utils"
	"net/http"
)

func Getdata(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, %s!\n", req.URL.Path[1:])
	collection := driver.ConnectBuilding()
	bname := utils.SetLookup(collection, "utility_bill", "bid")
	json.NewEncoder(res).Encode(bname)
}

func GetDB(res http.ResponseWriter, req *http.Request) {
	var building models.Building
	// , income, expenditure, land_cost, cost_of_land, utility_bill models.Building
	// var construction_cost_estimate, more_cost_expense, operating_cost models.Building
	// var interest_paid, corporate_income_tax models.Building
	var error models.Error

	json.NewDecoder(req.Body).Decode(&building)
	if building.Bid == "" {
		fmt.Println(building.Bid)
		error.Message = "Invalid Account"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(error)
		return
	} else {
		collection := driver.ConnectBuilding()
		bname := utils.SetLookup(collection, "utility_bill", "bid")
		json.NewEncoder(res).Encode(bname)
		// bname, error := buildingRepository.FindBuildngID(res, req, os.Getenv("building_name_collection"), building)
		// income, error := buildingRepository.FindBuildngID(res, req, os.Getenv("estimate_income_collection"), building)
		// expenditure, error := buildingRepository.FindBuildngID(res, req, os.Getenv("estimate_expenditure_collection"), building)
		// land_cost, error := buildingRepository.FindBuildngID(res, req, os.Getenv("land_cost_collection"), building)
		// cost_of_land, error := buildingRepository.FindBuildngID(res, req, os.Getenv("cost_of_land_collection"), building)
		// utility_bill, error := buildingRepository.FindBuildngID(res, req, os.Getenv("utility_bill_collection"), building)
		// construction_cost_estimate, error := buildingRepository.FindBuildngID(res, req, os.Getenv("construction_cost_estimate_collection"), building)
		// more_cost_expense, error := buildingRepository.FindBuildngID(res, req, os.Getenv("more_cost_expense_collection"), building)
		// operating_cost, error := buildingRepository.FindBuildngID(res, req, os.Getenv("operating_cost_collection"), building)
		// interest_paid, error := buildingRepository.FindBuildngID(res, req, os.Getenv("interest_paid_collection"), building)
		// corporate_income_tax, error := buildingRepository.FindBuildngID(res, req, os.Getenv("corporate_income_collection"), building)
		// building, error := buildingRepository.FindBuildngID(res, req, utility_bill_collection, building)
		if error.Message == "" {
			// json.NewEncoder(res).Encode(bname)
			// json.NewEncoder(res).Encode(income)
			// json.NewEncoder(res).Encode(expenditure)
			// json.NewEncoder(res).Encode(land_cost)
			// json.NewEncoder(res).Encode(cost_of_land)
			// json.NewEncoder(res).Encode(utility_bill)
			// json.NewEncoder(res).Encode(construction_cost_estimate)
			// json.NewEncoder(res).Encode(more_cost_expense)
			// json.NewEncoder(res).Encode(operating_cost)
			// json.NewEncoder(res).Encode(interest_paid)
			// json.NewEncoder(res).Encode(corporate_income_tax)
		}
	}
}

func GetJSON(response http.ResponseWriter, request *http.Request) {
	var building models.Building

	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &building)
	if err != nil {
		fmt.Println("error:", err)
	}
	json.NewEncoder(response).Encode(building)
}
