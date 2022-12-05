package main

import "fmt"

//var pl = fmt.Println()

// func main() {
// 	fmt.Println(" hi vishal")
// 	var arr = [4]int{1, 2, 3}
// 	//arr := {1,2,3,4}

// 	fmt.Println(arr)

// 	//var a = arr(append(4))

// func main() {

// 	arr1 := []int{1, 2, 3, 4}
// 	a := 0
// 	for i := 0; i < len(arr1); i++ {
// 		a += arr1[i]

// 	}
// 	fmt.Println(a)
// }

// func main() {
// 	arr1 := []int{1, 2, 3, 4}
// 	//s:=0
// 	for i, v := range arr1 {
// 		fmt.Println("index:", i, "Value:", v)
// 	}

// func main() {

// 	var a [5]int

// 	var i, j int
// 	for i = 0; i < 5; i++ {
// 		a[i] = i + 10
// 		//fmt.Println(i)

// 	}
// 	for j = 0; j < 5; j++ {
// 		fmt.Println(j, a[j])
// 	}
// }

// func main(){

// 	arr1:= [2][3]int {{1,2},{1,2}{

// 		fmt.Prinln(arr1)
// 	}
// }

func main() {
	// var s = []int{1, 2, 3, 4, 5}
	// fmt.Println(s)
	// s1 := append(s, 10)
	// fmt.Println(s1)
	// s2 := s[1:3]
	// fmt.Println(s2)
	var s = []string{"Vish", "Vinz", "Suhu", "Nandu"}

	fmt.Println(s)

	s1 := s[0:2]
	s2 := s[1:3]
	fmt.Println(s1, s2)
	s2[1] = "prom"
	s4 := append(s2, "moin")
	fmt.Println(s4)
}
