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
)

func Getdata(res http.ResponseWriter, req *http.Request) {
	collection := driver.ConnectBuilding()
	data := utils.GetBuildingByID(collection, "1234")
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
