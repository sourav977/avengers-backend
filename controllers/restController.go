package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sourav977/avengers-backend/helper"
	"github.com/sourav977/avengers-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllAvengers(w http.ResponseWriter, _ *http.Request) {
	var avengers []models.Avenger
	collection := helper.ConnectToDB()

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.SetError(err, http.StatusInternalServerError, w)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var avenger models.Avenger
		err := cur.Decode(&avenger)
		if err != nil {
			log.Fatal(err)
		}
		avengers = append(avengers, avenger)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(avengers)
}

func AddAvenger(w http.ResponseWriter, r *http.Request) {
	var avenger models.Avenger
	collection := helper.ConnectToDB()
	//decode req body into emp
	_ = json.NewDecoder(r.Body).Decode(&avenger)

	result, err := collection.InsertOne(context.TODO(), avenger)

	if err != nil {
		helper.SetError(err, http.StatusInternalServerError, w)
		return
	}
	//json returntype
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func Healthcheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Readiness(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
