package models

type Avenger struct {
	Name   string `json:"name" bson:"name,required"`
	Alias  string `json:"alias" bson:"alias,required"`
	Weapon string `json:"weapon" bson:"weapon,required"`
}
