package main

import "fmt"

func main() {
	fmt.Println("SLICE DECLERATION")
	var vinz1 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(vinz1)
	fmt.Println(len(vinz1), cap(vinz1))

	vinz1 = vinz1[0:5]
	fmt.Println(vinz1)
	fmt.Println(len(vinz1), cap(vinz1))

	vinz3 := vinz1[0:4]
	fmt.Println(vinz3)
	fmt.Println(len(vinz3), cap(vinz3))

	vinz2 := []struct {
		i int
		b bool
	}{{1, true}, {2, false}}
	fmt.Println(vinz2)

}
