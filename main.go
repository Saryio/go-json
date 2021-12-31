package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Person struct {
		Name  string `json:name`
		Email string `json:email`
	}

	p1 := Person{
		Name: "Samuel Antony",
		Email: "dev.saryio@gmail.com",
	}

	p2 := Person{
		Name: "Foo Foo",
		Email: "foo@foo.com",
	}

	persons := []Person{p1, p2}

	rawPersons, err := json.Marshal(persons)

	if err != nil{
		fmt.Println(err.Error())
	}

	fmt.Println("Antes:\n",persons)

	fmt.Println("\nDepois:\n", string(rawPersons))

}
