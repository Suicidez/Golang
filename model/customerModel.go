package model

import (
	"context"
	"ewallet/utill"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = utill.DBConnection().Collection("EW_CUSTOMER")

type CustomerModel struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FistName string             `bson:"firstname" json: firstname`
	LastName string             `bson:"lastname" json: lastname`
}

func (e CustomerModel) New(
	_id primitive.ObjectID,
	fistname string,
	lastname string,
) {
	e.ID = _id
	e.FistName = fistname
	e.LastName = lastname
}

func (e CustomerModel) Select() []*CustomerModel {
	cursor, err := collection.Find(
		context.TODO(),
		bson.M{},
		options.Find(),
	)
	if err != nil {
		log.Fatal(err)
	}

	var result []*CustomerModel
	for cursor.Next(context.TODO()) {
		var elem CustomerModel
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(elem)
		result = append(result, &elem)
	}
	cursor.Close(context.TODO())
	fmt.Println(*&result[0].ID)
	return result
}

func (e CustomerModel) Insert() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	insertResult, err := collection.InsertOne(ctx, e)
	defer cancel()
	if err != nil {
		log.Fatal((err))
	}
	fmt.Println(insertResult.InsertedID)
}

func (e CustomerModel) CheckData() {
	fmt.Println(e)
}
