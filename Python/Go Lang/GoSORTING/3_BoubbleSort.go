package main

import "fmt"

func BubbleSort(A []int) {
	n := len(A)
	for i := 0; i <= n; i++ {
		for j := 0; j < n-i-1; j++ {
			if A[j] > A[j+1] {
				temp := A[j]
				A[j] = A[j+1]
				A[j+1] = temp
			}

		}
	}
}
func main() {
	A := []int{2, 5, 4, 9, 1, 3}
	fmt.Println("Original Array:", A)
	BubbleSort(A)
	fmt.Println("Sorted Array: ", A)
}
