package main
import(
	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"strings"
	"io"
	"io/ioutil"
	"net/http"
	//"strconv"
)

type User struct{
	Id int32 `json:"id"`
	Name string `json:"name"`
}

func hello(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","text/json")
	io.WriteString(res,`{"id":1,"name":"Heena"}`)
}
func book (res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type","text/json")
	io.WriteString(res,`{"bookid":1001,"book":"god"}`)
}
func withget(res http.ResponseWriter,req *http.Request){

	res.Header().Set("Content-Type","text/json")
	response,_:=http.Get("http://jsonplaceholder.typicode.com/users")
    data,_:=ioutil.ReadAll(response.Body)
	var user []User

	json.NewDecoder(strings.NewReader(string(data))).Decode(&user)
	json.NewEncoder(res).Encode(user)


}
func main(){
	router := mux.NewRouter()
	router.HandleFunc("/hello",hello)
	router.HandleFunc("/book",book)
	router.HandleFunc("/withget",withget).Methods("Get")

	http.Handle("/",router)
	http.ListenAndServe(":8080",nil)
}
