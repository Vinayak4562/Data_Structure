package main

import "fmt"

type node struct {
	element int
	next    *node
}
type SinglyLL struct {
	head *node
	tail *node
	size int
}

func (l *SinglyLL) len() int {
	return l.size
}
func (l *SinglyLL) isempty() bool {
	return l.size == 0
}
func (l *SinglyLL) addfirst(e int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.head = newest
		l.tail = newest
	}
	newest.next = l.head
	l.head = newest
	l.size++
}
func (l *SinglyLL) addlast(e int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.head = newest
		l.tail = newest
	}
	p := l.head
	i := 1
	for i < l.len()-1 {
		p = p.next
		i++
	}
	l.tail.next = newest
	l.tail = newest
	l.size++
}
func (l *SinglyLL) addany(e, position int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.head = newest
		l.tail = newest
	}
	p := l.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	newest.next = p.next
	p.next = newest
	l.size++
}

func (l *SinglyLL) display() {
	p := l.head
	i := 0
	for i < l.len() {
		fmt.Print(p.element, "-->")
		p = p.next
		i++

	}
	fmt.Println()
}

func main() {
	a := &SinglyLL{}
	a.addfirst(10)
	a.addfirst(20)
	a.addfirst(30)
	a.display()
	fmt.Println(a.len())
	a.addlast(1)
	a.addlast(2)
	a.display()
	a.addany(222, 2)
	a.display()
	a.addany(222, 7)
	a.display()
}
