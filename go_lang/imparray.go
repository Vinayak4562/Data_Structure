package main

import "fmt"

func main() {
	fmt.Println("ARRAY DECLERATION")
	var vinz1 = [3]int{1, 2, 3}
	fmt.Println(vinz1)

	var vinz2 = [3]string{"Vinayak", "Vishal", "Suhas"}
	fmt.Println(vinz2)

	var vinz3 = [3]float64{0.1, 2.1, 3}
	fmt.Println(vinz3)

	// var vinz [5]int
	// var i, j int

	// for i = 0; i < len(vinz); i++ {
	// 	vinz[i] = i + 10
	// }
	// for j = 0; j < len(vinz); j++ {
	// 	fmt.Printf("Element[%d] = %d\n", j, vinz[j])
	// }

	//Multi dimensional arrays is simply a list of one-dimensional arrays

	var vin1 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(vin1)
	var i, j int
	for i = 0; i < 2; i++ {
		for j = 0; j < 3; j++ {
			fmt.Print(vin1[i][j])
		}
		fmt.Println()
	}

}
