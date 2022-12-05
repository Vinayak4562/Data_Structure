package main

import "fmt"

var a = 4 // Globale or main variable

func demo() {
	var a = 9 // Function or local variable
	fmt.Println(a)
}

func main() {
	demo()
	fmt.Println(a)
}
