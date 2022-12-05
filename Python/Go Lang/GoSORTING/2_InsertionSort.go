package main

import "fmt"

func InsertionSort(A []int) {

	for i := 1; i < len(A); i++ {
		Cvalue := A[i]
		position := i
		for position > 0 && A[position-1] > Cvalue {
			A[position] = A[position-1]
			position = position - 1
		}
		A[position] = Cvalue
	}

}

func main() {
	A := []int{3, 5, 8, 9, 6, 2}
	fmt.Println("Original Arraay: ", A)
	InsertionSort(A)
	fmt.Println("Sorted Arraay: ", A)
}
