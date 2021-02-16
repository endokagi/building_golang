package models

type UserDocker struct {
	Firstname string `json:"firstname, omitempty" bson:"firstname, omitempty"`
	Lastname  string `json:"lastname, omitempty" bson:"lastname, omitempty"`
	Email     string `json:"email, omitempty" bson:"email, omitempty"`
	Password  string `json:"password, omitempty" bson:"password, omitempty"`
}
