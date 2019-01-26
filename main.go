package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*Models*/

//Person Struck: is the strcut in which we define the person properties.
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

//Address Struck: is where we define the location of the person.
type Address struct {
	City  string `json:"address,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

/*Controllers*/

//GetPeople function: sends all the people we have in data base.
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//GetPerson function: sends one person defined in the db according to the ID
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//CreatePerson function: creates a new person and adds it to the people db
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

//DeletePerson function: Erases a person.
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

/*Server code*/
func main() {
	/*Se instancia la funcionalidad de enrutamiento*/
	router := mux.NewRouter()

	/*Se definene las rutas de accesso*/
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	/*hardCoded data*/
	people = append(people, Person{ID: "1", FirstName: "Luis", Lastname: "Sandoval", Address: &Address{City: "Bogotá", State: "Bogotá"}})
	people = append(people, Person{ID: "2", FirstName: "Nataly", Lastname: "Camargo", Address: &Address{City: "Bogotá", State: "Bogotá"}})

	/*Se registra cualquier error que pueda suceder al levantar el servidor en el puerto 3000*/
	log.Fatal(http.ListenAndServe(":3000", router))
	fmt.Println("server running on port 3000")
}
