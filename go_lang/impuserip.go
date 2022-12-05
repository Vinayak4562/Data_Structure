package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var mystring string
	// fmt.Scanln(&mystring)
	// fmt.Println(mystring)

	// var name string
	// fmt.Scanln(&name)
	// var a, _ = fmt.Println(name)
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter Your Full name: ")
	// myname, _ := reader.ReadString('\n')
	// fmt.Println(myname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Ratings: ")
	myrating, _ := reader.ReadString('\n')
	mynumratings, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumratings)

}
