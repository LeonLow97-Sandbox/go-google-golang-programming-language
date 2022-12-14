package main

import "fmt"

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// automatically increments by 1
const (
	a = iota
	b
	c 
)

const (
	d = iota
	e
	f 
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("")

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}