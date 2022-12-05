package main

import "fmt"

func main() {
	a := 20
	b := 10

	if a == b {
		fmt.Println("a is equal to b")
	} else if a > b {
		fmt.Println("a is greater than to b")
	} else {
		fmt.Println("a is less than to b")
	}
}
