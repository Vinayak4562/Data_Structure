package main

import "fmt"

func main() {
	vinz := map[int]string{1: "vinz", 2: "promz", 3: "Ayush"}
	// var vinz map[string]int{"vinz": 1, "promz": 2, "Ayush": 3}

	vinz[3] = "vinod"
	vinz[4] = "vinayak"
	vinz[7] = "pramod"
	vinz[5] = "pammi"
	vinz[6] = "pammya"
	fmt.Println()
	fmt.Println(vinz)

	delete(vinz, 4)
	delete(vinz, 2)
	fmt.Println(vinz)

	fmt.Println(vinz[3])

	i, ok := vinz[5]
	fmt.Println("The value :", i, "Present ?", ok)
	j, ok := vinz[4]
	fmt.Println("The value :", j, "Present ?", ok)
}
