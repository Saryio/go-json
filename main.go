package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name  string `json:name`
	Email string `json:email`
}

func main() {

	p1 := Person{Name: "Samuel Antony", Email: "dev.saryio@gmail.com"}
	p2 := Person{Name: "Fernanda Elloise", Email: "fefa@gmail.com"}

	persons := []Person{p1, p2}

	arrPerson, err := json.Marshal(persons)

	if err != nil {
		log.Println(err)
	}

	log.Println(string(arrPerson))
}
