package main

import "fmt"

// creating node
type Node struct {
	data int
	next *Node
	prev *Node
}

// Creating doubly linked list
type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

// creating lenth method for counting the nodes
func (dl *DoublyLinkedList) len() int {
	return dl.size
}

//  creating isempty method for checking that list is empty or not
func (dl *DoublyLinkedList) isempty() bool {
	return dl.size == 0
}

// adding node at the First in the list using addfirst method
func (dl *DoublyLinkedList) addfirst(n int) {
	new := &Node{n, nil, nil}
	if dl.isempty() {
		// dl.head = new
		dl.tail = new
	} else {
		dl.head.prev = new
		new.next = dl.head
		// dl.head = new
	}
	dl.head = new
	dl.size++
}

// adding node at the last in the list using addlast method
func (dl *DoublyLinkedList) addlast(n int) {
	new := &Node{n, nil, nil}
	if dl.isempty() {
		dl.head = new
		// dl.tail = new
	} else {
		dl.tail.next = new
		new.prev = dl.tail
		// dl.tail = new
	}
	dl.tail = new
	dl.size++
}

// Adding node anywhere in the list using specifide index position
func (dl *DoublyLinkedList) addany(n int, index int) {
	if index == 0 {
		dl.addfirst(n)
	} else if index >= (dl.len()) {
		dl.addlast(n)
	} else {
		new := &Node{n, nil, nil}
		p := dl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		new.prev = p
		new.next = p.next
		p.next.prev = new
		p.next = new
	}
	dl.size++
}

// removing the fisrt node in the list using remvfirst method
func (dl *DoublyLinkedList) removefirst() {
	if dl.isempty() {
		fmt.Println("Doubly Linked List is empty")
	} else {
		dl.head = dl.head.next
		dl.size--
	}
	if dl.isempty() {
		dl.head = nil
		dl.tail = nil
	}
}

// removing the Last node in the list using remvlast method
func (dl *DoublyLinkedList) removelast() {
	if dl.isempty() {
		fmt.Println("Doubly linked list is empty")
	} else {
		dl.tail = dl.tail.prev
		dl.tail.next = nil
		dl.size--
	}
	if dl.isempty() {
		dl.head = nil
		dl.tail = nil
	}
}

// removing node anywhere in the list using specifide index possion
func (dl *DoublyLinkedList) removeany(index int) {
	if index >= dl.len() {
		fmt.Println("Index out of range")
	} else if index == 0 {
		dl.removefirst()
	} else if index == (dl.len())-1 {
		dl.removelast()
	} else {
		p := dl.head
		i := 1
		for i < index {
			p = p.next
			i++
		}
		p.next = p.next.next
		p.next.prev = p
		dl.size--
	}
	if dl.isempty() {
		dl.head = nil
		dl.tail = nil
	}
}

//  creating the display function
func (dl *DoublyLinkedList) display(n int) {
	if n == 0 {
		p := dl.head
		for p != nil {
			fmt.Print(p.data, "-->")
			p = p.next
		}
		fmt.Println()
	} else if n == 1 {
		p := dl.tail
		for p != nil {
			fmt.Print(p.data, "-->")
			p = p.prev
		}
		fmt.Println()
	} else {
		fmt.Println("Invalid argument: arugument must be 0 for forword display or 1 for reverse display")
	}
}

//  Main function for creating objects a clling them
func main() {
	d := DoublyLinkedList{}
	d.addfirst(40)
	d.addfirst(30)
	d.addfirst(20)
	d.display(10)
	d.display(1)
	// 	d.addlast(40)
	// 	d.addlast(50)
	// 	d.addlast(60)
	// 	d.addany(11, 2)
	// 	d.display(0) // 0 --> forword diplay, 1--> reverse display
	// 	d.removefirst()
	// 	d.removefirst()
	// 	d.display(0)
	// 	d.removelast()
	// 	d.removelast()
	// 	d.display(0)
	// 	d.removeany(d.len() - 1)
	// 	d.removeany(3)
	// 	d.display(0)
	// 	fmt.Println(d.len())
	// d.addfirst(10)
	// d.addfirst(20)
	// d.display(0)
	// d.removeany(0)
	// d.display(0)
}
