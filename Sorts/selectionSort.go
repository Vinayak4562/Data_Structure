package main

import "fmt"

func selectionsort(A []int) {
	for i := 0; i < len(A)-1; i++ {
		position := i
		for j := i + 1; j < len(A); j++ {
			if A[position] > A[j] {
				position = j
			}
		}
		temp := A[i]
		A[i] = A[position]
		A[position] = temp
	}
}

func main() {
	A := []int{3, 5, 8, 9, 6, 2}
	fmt.Println("Original Array: ", A)
	selectionsort(A)
	fmt.Println("Array after sorting: ", A)
}
