package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}

func foo(w http.ResponseWriter, _ *http.Request) {
	p1 := person{
		First: "Jenny",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data", err)
	}
}

func bar(_ http.ResponseWriter, r *http.Request) {
	var p1 person

	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data", err)
	}

	log.Println("Person", p1)
}
