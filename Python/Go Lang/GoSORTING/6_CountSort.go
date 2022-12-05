package main

import "fmt"

func findmax(arr []int) int {
	max_num := arr[0]
	for _, elem := range arr {
		if max_num < elem {
			max_num = elem
		}
	}
	return max_num
}
func CountSort(arr []int) {
	n := len(arr)
	m := findmax(arr)
	carray := make([]int, (m + 1))
	for i := 0; i < n; i++ {
		carray[arr[i]] = carray[arr[i]] + 1

	}
	i := 0
	j := 0
	for i = 0; i < m+1; {
		if carray[i] > 0 {
			arr[j] = i
			j++
			carray[i] = carray[i] - 1
		} else {
			i++
		}
	}

}

func main() {
	arr := []int{7, 1, 5, 2, 2}
	fmt.Println("Original Array", arr)
	CountSort(arr)
	fmt.Println("orted array", arr)
}
