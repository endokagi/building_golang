package buildingRepository

import (
	"context"
	"encoding/json"
	"jwt/driver"
	"jwt/models"
	"net/http"

	"github.com/subosito/gotenv"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	gotenv.Load()
}

func BuildingID(response http.ResponseWriter, request *http.Request, bname models.BuildingName) (string, models.Error) {
	var collection = driver.ConnectBuilding()
	var building models.BuildingName = bname
	var error models.Error
	err := collection.FindOne(context.TODO(), bson.M{"bid": building.Bid}).Decode(&building)
	if err != nil {
		error.Message = "Invalid Account"
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(error)

	}
	return building.Bid, error
}
