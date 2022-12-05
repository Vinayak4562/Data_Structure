package main

import "fmt"

func main() {
	var a int = 7
	fmt.Println(a)

	fmt.Printf("%T\n", a)
	b := []int{1, 2, 3}
	fmt.Println(b[1])
	b = append(b, 4, 5, 6)
	fmt.Println(b)
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

}
