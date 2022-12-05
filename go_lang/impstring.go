package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println
	p("STRING DECLERATION")

	vinz := strings.Contains("HI good morning every one", "one")
	p(vinz)

	vinz1 := strings.Split("a,b,c", ",")
	// fmt.Printf("%T\n", vinz1)
	fmt.Printf("%#v\n", vinz1)

	vinz3 := strings.Split("HI good morning every one", "e")
	// fmt.Printf("%T\n", vinz3)
	fmt.Printf("%#v\n", vinz3)

	vinz4 := []string{"HI", "good", "morning", "every", "one"}
	s := strings.Join(vinz4, "-")
	// fmt.Printf("%#v", vinz3)
	p(s)
}
