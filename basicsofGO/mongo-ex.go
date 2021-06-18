package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testRecord struct {
	Name    string `bson:"name"`
	Address string `bson:"address"`
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://vrk:vrk@baprojects.bassure.in/vrk"))

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	testCollection := client.Database("vrk").Collection("test")

	var temp testRecord

	fmt.Print("Enter Name: ")
	fmt.Scanf("%s", &temp.Name)
	fmt.Scanf("%s", &temp.Address)
	fmt.Print("Enter Address: ")
	fmt.Scanf("%s", &temp.Address)

	testCollection.InsertOne(context.TODO(), temp)
}

func main1() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://vrk:vrk@baprojects.bassure.in/vrk"))

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	testCollection := client.Database("vrk").Collection("test")

	cursor, err := testCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var result testRecord
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
