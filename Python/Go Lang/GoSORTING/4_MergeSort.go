package main

import "fmt"

func mergesort(A []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergesort(A, left, mid)
		mergesort(A, mid+1, right)
		merge(A, left, mid, right)
	}
}

func merge(A []int, left, mid, right int) {
	i := left
	j := mid + 1
	k := left
	B := make([]int, (right + 1))

	for i <= mid && j <= right {
		if A[i] < A[j] {
			B[k] = A[i]
			i++
		} else {
			B[k] = A[j]
			j++
		}
		k++
	}
	for i <= mid {
		B[k] = A[i]
		i++
		k++
	}
	for j <= right {
		B[k] = A[j]
		j++
		k++
	}
	for x := left; x < right+1; x++ {
		A[x] = B[x]
	}

}

func main() {
	A := []int{2, 5, 8, 6, 4, 1}
	fmt.Println("Original Array:", A)
	mergesort(A, 0, len(A)-1)
	fmt.Println("Sorted Array:", A)
}
