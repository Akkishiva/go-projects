package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    http.HandleFunc("/person", getPerson)
    http.ListenAndServe(":8080", nil)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
    // Create a new Person struct and populate it with some data.
    person := Person{Name: "Alice", Age: 25}

    // Encode the Person struct as JSON and write it to the ResponseWriter.
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(person)
}
