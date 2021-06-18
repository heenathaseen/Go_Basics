package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	// "io"
	"bufio"
	"os"

	// "io/ioutil"
	// "log"
	"net/http"
	// "strings"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type COLLEGE struct {
	College_code int32  `bson:"collegecode"`
	College_name string `bson:"collegename"`
	Pin          string `bson:"pin"`
	State        string `bson:"state"`
	City         string `bson:"city"`
}

func main() {
	MenuItems()

}

func MenuItems() {
	var Option int32
	fmt.Println("*****************************")
	fmt.Println(".....COLLEGE DETAILS.....")
	fmt.Println("*****************************")
	fmt.Println(" 1.view colleges\n 2.delete college\n 3.Add college\n 4.exitprogram\n ")
	fmt.Printf("choose any one functionality:")
	fmt.Scanf("%d", &Option)
	SelectedOptions(Option)

}

func SelectedOptions(option int32) {
	switch option {
	case 1:
		findAll()
	case 2:
		deleteOne()
	case 3:
		addCollegesservice()
	case 4:
		fmt.Println("Thank you")
		os.Exit(0)
	default:
		fmt.Println("enter valid option")
		// fmt.Println("enter college code:")
		a := inputFromUsermongo()
		SelectedOptions(a)

	}
}

func addCollegesservice() {
	var clg COLLEGE
	fmt.Println("college code:")
	clg.College_code = inputFromUsermongo()
	fmt.Println("college name:")
	clg.College_name = inputFromUsermy()
	fmt.Println("pin:")
	clg.Pin = inputFromUsermy()
	fmt.Println("state:")
	clg.State = inputFromUsermy()
	fmt.Println("city:")
	clg.City = inputFromUsermy()
	fmt.Println(clg)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://uva:uva@baprojects.bassure.in:27017/uva"))
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("uva").Collection("colleges")
	_, err = collection.InsertOne(context.TODO(), clg)
	if err != nil {
		log.Fatal(err)
	}
	MenuItems()
}

func deleteOne() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://uva:uva@baprojects.bassure.in:27017/uva"))
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	fmt.Println("enter the college code that need to be deleted:")
	d := inputFromUsermongo()
	collection := client.Database("uva").Collection("colleges")
	_, err = collection.DeleteOne(context.TODO(), bson.D{{"collegecode", d}})
	MenuItems()
}

// func updateOne(){
// 	clinet, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://uva:uva@baprojects.bassure.in:27017/uva"))
// 	defer func() {
// 		if err = clinet.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	collection := clinet.Database("uva").Collection("colleges")

// 	err := collection.UpdateOne(context.TODO(),bson.D{{"collegecode", d}},bson.D{{$set,{"collegename","n"}}})


// }

func findAll() {
	fmt.Println("please enter your name")
	a := inputFromUsermy()
	fmt.Printf(" hi, %s welcome! ", a)
	clinet, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://uva:uva@baprojects.bassure.in:27017/uva"))
	defer func() {
		if err = clinet.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := clinet.Database("uva").Collection("colleges")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	defer cursor.Close(context.TODO())
	var clg1 COLLEGE
	var clg2 []COLLEGE

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&clg1)
		clg2 = append(clg2, clg1)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(clg1)
	}
	fmt.Println(clg2)
	MenuItems()
}

func httpHandling() {
	router := mux.NewRouter()
	router.HandleFunc("/add", collegeHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func inputFromUsermy() string {
	stdin := bufio.NewReader(os.Stdin)
	var a string
	for {
		_, err := fmt.Fscan(stdin, &a)
		if err == nil {
			break
		}
		stdin.ReadString('\n')
	}
	return a
}

func inputFromUsermongo() int32 {
	stdin := bufio.NewReader(os.Stdin)
	var a int32
	for {
		_, err := fmt.Fscan(stdin, &a)
		if err == nil {
			break
		}
		stdin.ReadString('\n')
	}
	return a
}

func collegeHandler(res http.ResponseWriter, r *http.Request) {
	var clg COLLEGE
	json.NewEncoder(res).Encode(clg)
}

func converting() {
	var clg COLLEGE
	s := `{"collegeCode":"44","collegeName":"mce","pin":"600000","state":"tn","city":"chennai"}`
	json.NewDecoder(strings.NewReader(s)).Decode(&clg)
	fmt.Println(clg)
}
