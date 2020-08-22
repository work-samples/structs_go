package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p1 := person{"James", "Url"}
	fmt.Printf("%+v\n", p1)

	p2 := person{firstName: "Henry", lastName: "Url"}
	fmt.Printf("%+v\n", p2)

	var p3 person
	fmt.Printf("%+v\n", p3)

	p3.firstName = "Daniel"
	p3.lastName = "Url"
	fmt.Printf("%+v\n", p3)
}
