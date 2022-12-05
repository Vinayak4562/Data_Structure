package main

import "fmt"

// Creating the node struct
type Node struct {
	data int
	next *Node
}

// Creating the link list struct
type linkedlist struct {
	head *Node
	tail *Node
	size int
}

// Creating the length function
func (l *linkedlist) len() int {
	return l.size
}

// Creating the isempty function
func (l *linkedlist) isempty() bool {
	return l.size == 0
}

// Creating the addfirst function for adding the element to the list
func (l *linkedlist) addfirst(e int) {
	newest := &Node{e, nil}
	if l.head == nil {
		l.head = newest
	} else {
		l.tail.next = newest
	}
	l.tail = newest
	l.size++
}

//  adding node at the last in the list using addlast
func (l *linkedlist) addlast(e int) {
	newest := &Node{e, nil}
	if l.isempty() {
		l.head = newest
	} else {
		l.tail.next = newest
	}
	l.tail = newest
	l.size++
}

func (l *linkedlist) addany(e int, position int) {
	newest := &Node{e, nil}
	p := l.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
		newest.next = p.next
		p.next = newest
		l.size++
	}
}

// Creating the display function for desplaying the element
func (l *linkedlist) disply() {
	p := l.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}

}

func main() {
	L := linkedlist{}
	L.addfirst(40)
	L.addfirst(30)
	L.addfirst(20)
	L.addfirst(10)
	L.disply()
	L.addlast(50)
	L.disply()

	L.addany(1, 3)
	L.disply()

}
