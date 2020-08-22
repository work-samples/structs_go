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

	var p3 person
	fmt.Println(p3)

	p3.firstName = "Daniel"
	fmt.Println(p3)

	p3.firstName = "Url"
	fmt.Println(p3)
}
