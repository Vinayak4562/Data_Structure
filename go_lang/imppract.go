package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the value: ")
	myinput, _ := reader.ReadString('\n')
	fmt.Println(myinput)
	// myinput2, _ := strconv.strings(myinput)

	// s := []string{myinput, myinput, "vinayak", "vishal", "moin"}
	fmt.Print(myinput)
	// fmt.Print(s)
	// s1 := myinput[1:]
	s2 := strings.Split(myinput, "")
	fmt.Printf("%T", s2)
	fmt.Println(len(myinput))
	// a := "ţară"
	a := "vinayak havannavar"
	b := utf8.RuneCountInString(a)
	fmt.Println(b)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%c\n", a[i])
	}

	// for i, j := range a {
	// 	fmt.Printf("i:%d, j:%c \n ", i, j)
	// }
	// fmt.Println(a[2:5])
	// fmt.Println(myinput[0])

}
