package utils

import (
	buildingRepository "bul/repository/building"

	"go.mongodb.org/mongo-driver/bson"
)

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
