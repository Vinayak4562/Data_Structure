package main

import "fmt"

func main() {
	fmt.Println("Welcome for Maps")
	//Initialising the Map
	languages := make(map[string]string)

	// Adding the values to the map
	languages["PY"] = "Python"
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"

	//printing the content of map
	fmt.Println("List of all languages: ", languages)

	//Getting the value using key
	fmt.Println("JS is short for ", languages["JS"])
	fmt.Println("PY is short for ", languages["PY"])
}
