package main
import(
	"strconv"
	"fmt"
	//"os"
	"log"
	"context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type studentRegistration struct{
	Regno int ` json:regno bson:regno `
	Name string `json:"name" bson:"name" `
	Dept string `json:"dept" bson:"dept" `
	Collegename string `json:"collegename" bson:"collegename" `
}
//var s studentRegistration


func main(){


router := mux.NewRouter()
router.HandleFunc("/postdetail", postDetailHandler).Methods("POST")
router.HandleFunc("/getalldetails", getallDetailsHandler).Methods("Get")
router.HandleFunc("/getdetail/{regno}", getDetail).Methods("Get")
router.HandleFunc("/deletedetail/{regno}", deleteDetail).Methods("Delete")
router.HandleFunc("/updatedetail/{regno}", updateDetail).Methods("Put")
http.Handle("/", router)
err := http.ListenAndServe(":8080", nil)
if err != nil {
	log.Fatal(err)
}


}

func postDetailHandler(res http.ResponseWriter, req *http.Request){
    var s studentRegistration

	res.Header().Set("Content-Type", "application/json")
	json.NewDecoder(req.Body).Decode(&s)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://heena:heena@baprojects.bassure.in:27017/heena"))
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("heena").Collection("studentregistration")
	_, err = collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(res).Encode(s)

}

func getallDetailsHandler(res http.ResponseWriter, req *http.Request){
    var s studentRegistration
	res.Header().Set("Content-Type", "application/json")
	// json.NewDecoder(req.Body).Decode(&person)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://heena:heena@baprojects.bassure.in:27017/heena"))
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("heena").Collection("studentregistration")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	defer cursor.Close(context.TODO())
	var students []studentRegistration
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&s)
		students = append(students, s)
	}
	json.NewEncoder(res).Encode(students)
}
func  getDetail(res http.ResponseWriter, req *http.Request){
	param := mux.Vars(req)
	d,_ := strconv.Atoi( param["regno"])
	//d := param["regno"]
	
	fmt.Println(d)
    var s studentRegistration

	res.Header().Set("Content-Type", "application/json")
	// json.NewDecoder(req.Body).Decode(&person)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://heena:heena@baprojects.bassure.in:27017/heena"))
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("heena").Collection("studentregistration")
	collection.FindOne(context.TODO(), bson.M{"regno": d}).Decode(&s)
	json.NewEncoder(res).Encode(s)
}
func deleteDetail(res http.ResponseWriter, req *http.Request){
	param := mux.Vars(req)
	var s studentRegistration
	d,_ := strconv.Atoi( param["regno"])

	res.Header().Set("Content-Type", "application/json")
	json.NewDecoder(req.Body).Decode(&s)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://heena:heena@baprojects.bassure.in:27017/heena"))
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("heena").Collection("studentregistration")
	collection.DeleteOne(context.TODO(), bson.D{{"regno", d}})

}

func updateDetail(res http.ResponseWriter, req *http.Request){
	param := mux.Vars(req)
	d,_ := strconv.Atoi( param["regno"])

	fmt.Println(d)
	var s studentRegistration

	res.Header().Set("Content-Type", "application/json")
	json.NewDecoder(req.Body).Decode(&s)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://heena:heena@baprojects.bassure.in:27017/heena"))
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("heena").Collection("studentregistration")
	collection.FindOneAndUpdate(context.TODO(), bson.M{"regno": d}, bson.D{
		{"$set", bson.D{
			{"regno",s.Regno},
			{"name", s.Name},
			{"dept", s.Dept},
			{"collegename", s.Collegename},
		}},
	})
	json.NewEncoder(res).Encode(s)

}





