package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	//s2 := s1
	//fmt.Println("original slice: ", s1, "\n", "Copied Slice: ", s2)
	//s2 := s1[0:3]
	//s2 := s1[0:]
	//s2 := s1[:len(s1)]
	s2 := s1[:3]

	fmt.Println(s2)
}
