package main

import "fmt"

func add(x int, y int) int {
	var out = x + y
	return out

}
func sub(x int, y int) int {
	var out = x - y
	return out

}

func mul(x int, y int) int {
	var out = x * y
	return out

}

func div(x int, y int) int {
	var out = x / y
	return out

}

func main() {

	num1 := 5
	num2 := 4
	result := add(num1, num2)
	result1 := sub(num1, num2)
	result2 := mul(num1, num2)
	result3 := div(num1, num2)
	fmt.Println("addition is:", result, "substraction is:", result1, "multiplication is:", result2, "Division is:", result3)
}
