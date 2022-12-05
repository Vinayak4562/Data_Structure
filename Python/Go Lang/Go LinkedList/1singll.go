package main

import "fmt"

// Creating Node
type Node struct {
	data int
	next *Node
}

// Creating  Singly Linked List
type SinglyLinkedList struct {
	head *Node
	tail *Node
	size int
}

// Creating length method for knowing the how many nodes are in list
func (sl *SinglyLinkedList) len() int {
	return sl.size
}

//  Searching element in the list
func (sl *SinglyLinkedList) search(key int) int {
	p := sl.head
	index := 0
	for p != nil {
		if p.data == key {
			return index
		}
		p = p.next
		index++
	}
	return -1
}

//  Adding first node in the list
func (sl *SinglyLinkedList) addfirst(n int) {
	new := &Node{n, nil}
	if sl.head == nil {
		sl.head = new
		sl.tail = new
	} else {
		new.next = sl.head
		sl.head = new
	}
	sl.size++
}

//  Adding last node to the list
func (sl *SinglyLinkedList) addlast(n int) {
	new := &Node{n, nil}
	if sl.head == nil {
		sl.head = new
	} else {
		sl.tail.next = new
	}
	sl.tail = new
	sl.size++
}

// adding anywher in the list by passing index value
func (sl *SinglyLinkedList) addany(n int, index int) {
	new := &Node{n, nil}
	p := sl.head
	count := 1
	for count < index {
		p = p.next
		count++
	}
	new.next = p.next
	p.next = new
	sl.size++
}

// Removeing first node from the list
func (sl *SinglyLinkedList) removefirst() {
	if sl.head == nil {
		fmt.Println("Linked list is empty")
	} else {
		p := sl.head.next
		sl.head = p
		sl.size--
	}
	if sl.head == nil {
		sl.tail = nil
	}
}

// Removeing last node from the list
func (sl *SinglyLinkedList) removelast() {
	if sl.head == nil {
		fmt.Println("Linked list is empty")
	} else {
		p := sl.head
		e := 1
		for e < (sl.len())-1 {
			p = p.next
			e++
		}
		sl.tail = p
		sl.tail.next = nil
		sl.size--
	}
}

// Removeing node Any where in the list
func (sl *SinglyLinkedList) removeany(i int) {

	if i != 0 && i != (sl.len())-1 && i < sl.len() {
		if sl.head == nil {
			fmt.Println("Linked list is empty")
		} else {
			p := sl.head
			count := 1
			for count < i {
				p = p.next
				count++
			}
			p.next = p.next.next
			sl.size--
		}
	} else if i == 0 {
		sl.removefirst()
	} else if i == (sl.len())-1 {
		sl.removelast()
	} else {
		fmt.Println("Index ot of range...!")
	}

}

//  For displying the list by traversing the list
func (sl *SinglyLinkedList) display() {
	p := sl.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
}

func main() {
	l := SinglyLinkedList{}
	l.addfirst(10)
	l.addfirst(20)
	l.addfirst(30)
	l.display()
	fmt.Println()

	l.addlast(40)
	l.addlast(50)
	l.addlast(60)
	l.display()
	fmt.Println()

	// l.addany(5, 2)
	// l.display()
	// fmt.Println()

	l.removefirst()
	l.display()
	fmt.Println()

	l.removelast()
	fmt.Println(l.len())
	l.display()
	fmt.Println()

	// l.removeany(2)
	// l.display()
	// fmt.Println()
	// fmt.Println(l.len())

	// l.addany(5, 10)
	// l.display()
	// fmt.Println()

}
