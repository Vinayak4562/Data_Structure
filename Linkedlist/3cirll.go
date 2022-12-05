package main

import "fmt"

// Creating Node
type Node struct {
	data int
	next *Node
}

// Creating  Circular Linked List
type CircularLinkedList struct {
	head *Node
	tail *Node
	size int
}

// Creating length method for knowing the how many nodes are in list
func (cl *CircularLinkedList) len() int {
	return cl.size
}

// Creating length method for knowing the how many nodes are in list
func (cl *CircularLinkedList) isempty() bool {
	return cl.size == 0
}

//  Searching element in the list
func (cl *CircularLinkedList) search(v int) int {
	p := cl.head
	i := 0
	for i < cl.len() {
		if p.data == v {
			return i
		}
		p = p.next
		i++
	}
	return -1
}

//  Adding first node in the list
func (cl *CircularLinkedList) addfirst(n int) {
	new := &Node{n, nil}
	if cl.isempty() {
		new.next = new
		cl.tail = new
	} else {
		new.next = cl.head
	}
	cl.head = new
	cl.size++
}

//  Adding last node to the list
func (cl *CircularLinkedList) addlast(n int) {
	new := &Node{n, nil}
	if cl.isempty() {
		new.next = new
		cl.head = new
	} else {
		new.next = cl.head
		cl.tail.next = new
	}
	cl.tail = new
	cl.size++
}

// adding anywher in the list by passing index value
func (cl *CircularLinkedList) addany(n int, index int) {
	if index == 0 {
		cl.addfirst(n)
	} else if index > (cl.len())-1 {
		cl.addlast(n)
	} else {
		new := &Node{n, nil}
		p := cl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		new.next = p.next
		p.next = new
		cl.size++
	}
}

// Removeing first node from the list
func (cl *CircularLinkedList) removefirst() {
	if cl.isempty() {
		fmt.Println("Circular Linked LIst is empty")
	} else {
		cl.head = cl.head.next
		cl.tail.next = cl.head
		cl.size--
	}
	if cl.isempty() {
		cl.head = nil
		cl.tail = nil
	}
}

// Removeing Last node from the list
func (cl *CircularLinkedList) removelast() {
	if cl.isempty() {
		fmt.Println("Circular Linked lis is empty")
	} else {
		p := cl.head
		i := 1
		for i < (cl.len() - 1) {
			p = p.next
			i++
		}
		p.next = cl.tail.next
		cl.tail = p
		cl.size--
	}
	if cl.isempty() {
		cl.head = nil
		cl.tail = nil
	}
}

// Removeing node Any where in the list
func (cl *CircularLinkedList) removeany(index int) {
	if index > (cl.len())-1 {
		fmt.Println("Index out of range")
	} else if index == (cl.len())-1 {
		cl.removelast()
	} else if index == 0 {
		cl.removefirst()
	} else {
		p := cl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		p.next = p.next.next
		cl.size--
	}
	if cl.isempty() {
		cl.head = nil
		cl.tail = nil
	}
}

//  For displying the list by traversing the list
func (cl *CircularLinkedList) display() {
	p := cl.head
	i := 0
	for i < cl.len() {
		fmt.Print(p.data, "-->")
		p = p.next
		i++
	}
	fmt.Println()
}

func main() {
	c := CircularLinkedList{}
	c.addfirst(20)
	c.addfirst(10)
	c.addfirst(00)
	c.display()

	c.addlast(30)
	c.addlast(40)
	c.display()

	c.addany(15, 2)
	c.display()

	fmt.Println(c.search(100))
	fmt.Println(c.search(30))

	c.removefirst()
	c.display()
	fmt.Println(c.len())

	c.removelast()
	c.display()

	c.removeany(2)
	c.display()

	// c.removeany(c.len())
	fmt.Println(c.len())

}
