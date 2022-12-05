package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome To Slices")
	var frutList = []string{"Apple", "Banana", "Mango", "Peach"}
	fmt.Println("Type of fruit list is : ", frutList) //O/P [Apple Banana Mango Peach]

	// If we want to add the element in am array we use append() it will automatically allocate the memory
	frutList = append(frutList, "Greps", "Pappaya")
	fmt.Println("Fruit list after appending : ", frutList) //O/P [Apple Banana Mango Peach Greps Pappaya]

	// WE can delete the element by using the append method by assigning the periculer index location
	// by remove the comma and assign the index valuve in [],it will star from possion number 1 it will count 1 2 3.. like this.
	frutList = append(frutList[1:])
	fmt.Println("Fruit list after removing[1:] : ", frutList) // O/P [Banana Mango Peach Greps Pappaya]

	frutList = append(frutList[1:3])
	fmt.Println("Fruit list after removing[1:3] : ", frutList) //Stars with 1 (0 is exclusive) and ends with 2 3rd one is exclusive
	// range are always non inclusive if we are dealings with a database it will help us a lot.
	//[Banana Mango Peach Greps Pappaya] O/P [ Mango Peach ]

	frutList = append(frutList[:3]) //[ Mango Peach ] O/P [Mango Peach Greps] it will add last valuve that is 3rd index
	fmt.Println("Fruit list after removing[:3] : ", frutList)

	/*highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 120
	highScores[2] = 640
	highScores[3] = 150

	fmt.Println(highScores)
	highScores = append(highScores, 666, 777, 111)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	//  we can remove the element on list using append highScores=[111 120 150 234 640 666 777]
	var index int = 4
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)
	*/
}
