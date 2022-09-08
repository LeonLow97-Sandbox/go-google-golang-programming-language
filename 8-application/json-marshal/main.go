package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Leon",
		last:  "Low",
		age:   25,
	}
	p2 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	// Marshal
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}