package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
)

type User struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street      string `json:"street"`
		City        string `json:"city"`
		Zipcode     string `json:"zipcode"`
		GeoLocation struct {
			Latitude  string `json:"lat"`
			Longitude string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
}

func webHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/json")
	io.WriteString(res,
		`{"id": 1234, "name": "John", "address": "U.S.A."}`)
}

func bookHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/json")
	io.WriteString(res,
		`{"id": 590, "title": "Go, How To Program", "publisher": "Unknown"}`)
}

func getUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/json")
	response, _ := http.Get("http://jsonplaceholder.typicode.com/users")

	data, _ := ioutil.ReadAll(response.Body)
	//io.WriteString(res, string(data))

	var currentUser []User
	json.NewDecoder(strings.NewReader(string(data))).Decode(&currentUser)
	json.NewEncoder(res).Encode(currentUser)
}

func getUser(res http.ResponseWriter, req *http.Request) {
	values := mux.Vars(req)

	res.Header().Set("Content-Type", "text/json")
	response, _ := http.Get("http://jsonplaceholder.typicode.com/users/" + values["user-id"])

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Fprintf(res, "[")
	io.WriteString(res, string(data))

	var currentUser User
	json.NewDecoder(strings.NewReader(string(data))).Decode(&currentUser)
	//json.NewDecoder(response.Body).Decode(&currentUser)

	//fmt.Fprintf(res, "%d # %s # %s # %s\n\n", currentUser.Id, currentUser.Name,
	//	currentUser.Username, currentUser.Email)
	fmt.Fprintf(res, ",")
	json.NewEncoder(res).Encode(currentUser)
	fmt.Fprintf(res, "]")
}

func getUsersWithCount(res http.ResponseWriter, req *http.Request) {
	values := mux.Vars(req)
	res.Header().Set("Content-Type", "text/json")
	response, _ := http.Get("http://jsonplaceholder.typicode.com/users")

	data, _ := ioutil.ReadAll(response.Body)
	//io.WriteString(res, string(data))

	var currentUsers []User
	json.NewDecoder(strings.NewReader(string(data))).Decode(&currentUsers)

	fromUserId, _ := strconv.Atoi(values["from-user-id"])
	count, _ := strconv.Atoi(values["count"])
	currentUsers = currentUsers[fromUserId-1 : fromUserId+count-1]
	json.NewEncoder(res).Encode(currentUsers)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", webHandler)
	router.HandleFunc("/book", bookHandler)
	router.HandleFunc("/users", getUsers).Methods("Get")
	router.HandleFunc("/users/{user-id}", getUser).Methods("Get")
	router.HandleFunc("/users/{from-user-id}/{count}", getUsersWithCount).Methods("Get")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
