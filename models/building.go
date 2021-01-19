package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuildingName struct {
	ID    primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid   string             `json:"bid, omitempty" bson:"bid, omitempty"`
	BName string             `json:"bname, omitempty" bson:"bname, omitempty"`
}

type EstimateIncome struct {
	ID   primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid  string             `json:"bid, omitempty" bson:"bid, omitempty"`
	Area []struct {
		AreaID primitive.ObjectID `json:"area_id, omitempty" bson:"area_id, omitempty"`
		Title  string             `json:"title, omitempty" bson:"title, omitempty"`
		Value  int                `json:"value, omitempty" bson:"value, omitempty"`
	}
}

type LandCost struct {
	ID               primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid              string             `json:"bid, omitempty" bson:"bid, omitempty"`
	Title            string             `json:"title, omitempty" bson:"title, omitempty"`
	N_land           int                `json:"nland, omitempty" bson:"nland, omitempty"`
	N_Perland        int                `json:"nPerLand, omitempty" bson:"nPerLand, omitempty"`
	TransterLandCost int                `json:"tranferLandCost, omitempty" bson:"tranferLandCost, omitempty"`
	MoreLandCost     []struct {
		Title string `json:"title, omitempty" bson:"title, omitempty"`
		Value int    `json:"value, omitempty" bson:"value, omitempty"`
		Text  string `json:"text, omitempty" bson:"text, omitempty"`
	}
}

type CostOfLand struct {
	ID         primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid        string             `json:"bid, omitempty" bson:"bid, omitempty"`
	Title      string             `json:"title, omitempty" bson:"title, omitempty"`
	CostOfLand int                `json:"costOfLand, omitempty" bson:"costOfLand, omitempty"`
}

type UtilityBill struct {
	ID          primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid         string             `json:"bid, omitempty" bson:"bid, omitempty"`
	Title       string             `json:"title, omitempty" bson:"title, omitempty"`
	UtilityBill []struct {
		Title string `json:"title, omitempty" bson:"title, omitempty"`
		Value int    `json:"value, omitempty" bson:"value, omitempty"`
		Text  string `json:"text, omitempty" bson:"text, omitempty"`
	}
}
type ConstrucEstimate struct {
	ID                primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid               string             `json:"bid, omitempty" bson:"bid, omitempty"`
	Title             string             `json:"title, omitempty" bson:"title, omitempty"`
	N_BuildingArea    int                `json:"nBuildingArea, omitempty" bson:"nBuildingArea, omitempty"`
	N_PerBuildingArea int                `json:"nPerBuildingArea, omitempty" bson:"nPerBuildingArea, omitempty"`
}

type MoreCostExpense struct {
	ID                   primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid                  string             `json:"bid, omitempty" bson:"bid, omitempty"`
	Title                string             `json:"title, omitempty" bson:"title, omitempty"`
	OwnershipTransferFee int                `json:"ownershipTransferFee, omitempty" bson:"ownershipTransferFee, omitempty"`
	MoreCostExpense      []struct {
		Title string `json:"title, omitempty" bson:"title, omitempty"`
		Value int    `json:"value, omitempty" bson:"value, omitempty"`
		Text  string `json:"text, omitempty" bson:"text, omitempty"`
	}
}

type OperatingCost struct {
	ID                primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid               string             `json:"bid, omitempty" bson:"bid, omitempty"`
	Title             string             `json:"title, omitempty" bson:"title, omitempty"`
	Commission        int                `json:"commission, omitempty" bson:"commission, omitempty"`
	CreditUsageFee    int                `json:"CreditUsageFee, omitempty" bson:"CreditUsageFee, omitempty"`
	MoreOperatingCost []struct {
		Title string `json:"title, omitempty" bson:"title, omitempty"`
		Value int    `json:"value, omitempty" bson:"value, omitempty"`
		Text  string `json:"text, omitempty" bson:"text, omitempty"`
	}
}

type InterestPaid struct {
	ID           primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid          string             `json:"bid, omitempty" bson:"bid, omitempty"`
	Title        string             `json:"title, omitempty" bson:"title, omitempty"`
	InterestPaid int                `json:"interestPaid, omitempty" bson:"interestPaid, omitempty"`
}

type CorporateIncomeTax struct {
	ID                 primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Bid                string             `json:"bid, omitempty" bson:"bid, omitempty"`
	Title              string             `json:"title, omitempty" bson:"title, omitempty"`
	CorporateIncomeTax int                `json:"corporateIncomeTax, omitempty" bson:"corporateIncomeTax, omitempty"`
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Building struct {
	Bid   string `json:"bid, omitempty" bson:"bid, omitempty"`
	BName string `json:"bname, omitempty" bson:"bname, omitempty"`

	Area []struct {
		Title string `json:"title, omitempty" bson:"title, omitempty"`
		Value int    `json:"value, omitempty" bson:"value, omitempty"`
	}

	N_land           int `json:"nland, omitempty" bson:"nland, omitempty"`
	N_Perland        int `json:"nPerLand, omitempty" bson:"nPerLand, omitempty"`
	TransterLandCost int `json:"tranferLandCost, omitempty" bson:"tranferLandCost, omitempty"`
	MoreLandCost     []struct {
		Title string `json:"title, omitempty" bson:"title, omitempty"`
		Value int    `json:"value, omitempty" bson:"value, omitempty"`
		Text  string `json:"text, omitempty" bson:"text, omitempty"`
	}

	CostOfLand int `json:"costOfLand, omitempty" bson:"costOfLand, omitempty"`

	UtilityBill []struct {
		Title string `json:"title, omitempty" bson:"title, omitempty"`
		Value int    `json:"value, omitempty" bson:"value, omitempty"`
		Text  string `json:"text, omitempty" bson:"text, omitempty"`
	}

	N_BuildingArea    int `json:"nBuildingArea, omitempty" bson:"nBuildingArea, omitempty"`
	N_PerBuildingArea int `json:"nPerBuildingArea, omitempty" bson:"nPerBuildingArea, omitempty"`

	OwnershipTransferFee int `json:"ownershipTransferFee, omitempty" bson:"ownershipTransferFee, omitempty"`
	MoreCostExpense      []struct {
		Title string `json:"title, omitempty" bson:"title, omitempty"`
		Value int    `json:"value, omitempty" bson:"value, omitempty"`
		Text  string `json:"text, omitempty" bson:"text, omitempty"`
	}

	Commission        int `json:"commission, omitempty" bson:"commission, omitempty"`
	CreditUsageFee    int `json:"CreditUsageFee, omitempty" bson:"CreditUsageFee, omitempty"`
	MoreOperatingCost []struct {
		Title string `json:"title, omitempty" bson:"title, omitempty"`
		Value int    `json:"value, omitempty" bson:"value, omitempty"`
		Text  string `json:"text, omitempty" bson:"text, omitempty"`
	}

	InterestPaid int `json:"interestPaid, omitempty" bson:"interestPaid, omitempty"`

	CorporateIncomeTax int `json:"corporateIncomeTax, omitempty" bson:"corporateIncomeTax, omitempty"`
}
