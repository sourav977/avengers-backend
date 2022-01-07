package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Avenger struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name,omitempty"`
	Alias  string             `json:"alias" bson:"alias,omitempty"`
	Weapon string             `json:"weapon" bson:"weapon,omitempty"`
}
