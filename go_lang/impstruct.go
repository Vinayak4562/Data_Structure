package main

import "fmt"

type vin struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	x := vin{firstName: "Vinayak", lastName: "Havannavr", age: 20}
	fmt.Println(x)
	fmt.Println(x.firstName)
	fmt.Println(x.age)
}
