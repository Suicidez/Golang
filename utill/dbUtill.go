package utill

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbConfig struct {
	Username     string
	Password     string
	DatabaseName string
}

func DBConnection() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbModel := dbConfig{
		Username:     "natdanai",
		Password:     "sifer007",
		DatabaseName: "ewallet",
	}
	dbURL := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.hdm95.mongodb.net/%s?retryWrites=true&w=majority",
		dbModel.Username, dbModel.Password, dbModel.DatabaseName)
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURL))
	err = client.Connect(ctx)
	db := client.Database("ewallet")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err, " Connect Fail")
	} else {
		fmt.Println("Connect Success")
	}
	return db
}
