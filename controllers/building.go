package controllers

import (
	"bul/models"
	"bul/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	buildingRepository "bul/repository/building"
)

func Getdata(res http.ResponseWriter, req *http.Request) {
	var bid = "0000"
	data := utils.GenerateData(bid)
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
		id, _ := buildingRepository.BuildingID(res, req, building.Bid)
		if id != building.Bid {
			error.Message = "Invalid ID"
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(error)
		} else {
			data := utils.GenerateData(id)
			json.NewEncoder(res).Encode(data)
		}

	}

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
