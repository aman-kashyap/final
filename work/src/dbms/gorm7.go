package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Person struct {
	gorm.Model
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	//Address   *Address `json:"address,omitempty"`
}

/*type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}*/

func init() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres port=5432 dbname=aman sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	r := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "ramu", Lastname: "kaka"}) //Address: &Address{City: "jammu", State: "j&k"}})
	people = append(people, Person{ID: "2", Firstname: "sonia", Lastname: "sharma"})
	r.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	r.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	r.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	r.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":12345", r))
}
