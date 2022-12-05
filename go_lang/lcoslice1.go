package main

import (
	"fmt"
)

func main() {

	cources := []string{"Java", "JavaScript", "ReactJS", "Shift", "Python", "Ruby"}
	fmt.Println("cources are listed here", cources)
	var index int = 3
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println("Ongoing Cources: ", cources)

}
