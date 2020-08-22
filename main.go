package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p1 := person{"James", "Url"}
	fmt.Println(p1)

	p2 := person{firstName: "Henry", lastName: "Url"}
	fmt.Println(p2)
}
