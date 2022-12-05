package main

import "fmt"

/*
func calc(x, y int) (int, int, int, int) {
	var add = x + y
	var sub = x - y
	var mul = x * y
	var div = x / y
	return add, sub, mul, div
}
*/

func calc(x, y int) (add, sub, mul, div int) {
	add = x + y
	sub = x - y
	mul = x * y
	div = x / y
	return
}

func main() {
	var num1 = 5
	var num2 = 4

	res1, res2, res3, res4 := calc(num1, num2)

	fmt.Println("addition is:", res1, "substraction is:", res2, "multiplication is:", res3, "Division is:", res4)
}
