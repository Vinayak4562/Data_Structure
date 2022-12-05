package main

import "fmt"

func main() {
	var vinz [3]string
	vinz[0] = "Vinayak"
	//vinz[1] = "Havannavar" here it will give a space at OP
	vinz[2] = "Pramod"
	//vinz[3]= "Ayush" it will through an error cz its created up to 1 2 3=3 memory location

	fmt.Println("Vinz containing the strings:", vinz)

	// also we can initializing the ARRAY like bellow
	vinz1 := [3]string{"Vinayak", "Havannavar", "Pramod"}
	fmt.Println("Vinz1 containing the strings:", vinz1)

	//Also we can iterate through for loop
	vinz2 := [4]int{1, 2, 3, 4}
	sum := 0
	for i := 0; i < len(vinz2); i++ {
		sum += vinz2[i]
	}
	fmt.Println("The sum of vinz is: ", sum)

	// we can use 2 Dimontional ARRAY
	vinz2DArray := [2][4]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(vinz2DArray)
	//also we can access the elements of an array
	fmt.Println(vinz2DArray[0][2])
	fmt.Println(vinz2DArray[1][2])
	fmt.Println(vinz2DArray[1][1])

}
