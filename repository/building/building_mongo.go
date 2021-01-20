package buildingRepository

import (
	"bul/driver"
	"bul/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/subosito/gotenv"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	gotenv.Load()
}

func BuildingID(res http.ResponseWriter, req *http.Request, bid string) (string, models.Error) {
	var collection = driver.ConnectBuilding()
	var building models.BuildingName
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var error models.Error
	err := collection.FindOne(ctx, bson.M{"bid": bid}).Decode(&building)
	if err != nil {
		error.Message = "Invalid ID"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(error)

	}
	return building.Bid, error
}
