package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter here: ")
	message, _ := reader.ReadString('\n')
	// mymessage, _ := strconv.ParseFloat(strings.TrimSpace(message), 64)
	fmt.Println(message)

	// s1 := []string{message}

	fmt.Print(message[0])
	// fmt.Print(s1)
	// s2 := strings.Split(message, "")
	b := utf8.RuneCountInString(message)
	fmt.Println(b)

	// fmt.Printf("%T", s2)
	// fmt.Println(len(message))

}
