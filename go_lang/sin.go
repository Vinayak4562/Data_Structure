package main

import (
	"fmt"
)

type stt struct {
	fname  string
	lname  string
	age    int
	salary float64
}

func myfunc(fname string, lname string, age int) {
	fmt.Println(fname, lname, age)
}

func main() {
	// e1 := stt{"vinayak", "havannavar", 25, 28950.50}
	// fmt.Println(e1)
	// fmt.Println(e1.fname)
	// fmt.Println(e1.lname)
	// fmt.Println(e1.age)
	// fmt.Println(e1.salary)
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your data here")
	// mydata, _ := reader.ReadString('\n')
	// fmt.Println(mydata)
	// // s := mydata[0]
	// // fmt.Println(s)
	// b := utf8.RuneCountInString(mydata)
	// fmt.Println(b)
	// for i := 0; i < len(mydata); i++ {
	// 	fmt.Printf("%c\n", mydata[i])

	// }
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter here")
	// mynum, _ := reader.ReadString('\n')
	// mydata, _ := strconv.ParseFloat(strings.TrimSpace(mynum), 64)
	// fmt.Print(mydata)

	myfunc("vinayak", "havannavar", 25)
	myfunc("vishal", "hirandagi", 25)

}
