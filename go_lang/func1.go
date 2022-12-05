package main

import "fmt"

func myfunc(x int, y string) (total int, resulr string) {
	total = x + x
	resulr = y + " world..!"
	return
}

func main() {
	a, b := myfunc(3, "hi")
	c, d := myfunc(4, "hello")

	fmt.Println(a, b)
	fmt.Println(c, d)
}
