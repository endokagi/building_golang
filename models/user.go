package models

import (
	// "time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	AccessToken  string             `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken string             `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty" bson:"refreshToken,omitempty"`
	AccessUuid   string `json:"accessUuid,omitempty" bson:"accessUuid,omitempty"`
	RefreshUuid  string `json:"refreshUuid,omitempty" bson:"refreshUuid,omitempty"`
	AtExpires    int64  `json:"atExpires,omitempty" bson:"atExpires,omitempty"`
	RtExpires    int64  `json:"rtExpires,omitempty" bson:"rtExpires,omitempty"`
}

type GetEmail struct {
	Email    string `json:"email"`
	Exp      int64  `json:"exp"`
	Password string `json:"password"`
}
