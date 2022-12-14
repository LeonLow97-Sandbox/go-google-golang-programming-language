package main

import (
	"errors"
	"log"
)

func main() {

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// // func New(text string) error)
		return 0, errors.New("square root of negative number")
	}
	return 42, nil
}
