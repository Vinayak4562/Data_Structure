package main

import "fmt"

func main() {

	// we have to use the var after criating it otherwise it will through an compile time error
	var x int = 7
	var y = 8
	z := 9
	var s1 string = "vinayak"
	s2 := "Havannavar"

	// we can mute the var by putting underscore "-". its act as a black hole.
	_ = s2

	fmt.Println("variable values ade:", x, y, z, "\n", s1)

	// Assigning multiple variable using " := "
	car, price := "BMW", 6000000

	// if we are assigning multiple values to this pridifine variabls using " :=" it will through an error
	// to avoid this we have to declear at least one new variable in tht group
	//car, price := "Audi",70000
	car2, price := "Audi", 70000
	fmt.Println(car, price, car2, price)

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	//fmt.Println(opened, file)

}
