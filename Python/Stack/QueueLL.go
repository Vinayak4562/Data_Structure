package main

import "fmt"

type node struct {
	element int
	next    *node
}
type Qlinkedlist struct {
	frount *node
	rear   *node
	size   int
}

func (l *Qlinkedlist) len() int {
	return l.size
}
func (l *Qlinkedlist) isempty() bool {
	return l.size == 0
}
func (l *Qlinkedlist) enqueue(e int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.frount = newest
	} else {
		l.rear.next = newest
	}
	l.rear = newest
	l.size++
}
func (l *Qlinkedlist) dqueue() {
	if l.isempty() {
		fmt.Println("Queue is empty")
		return
	}
	e := l.frount.element
	l.frount = l.frount.next
	l.size--
	if l.isempty() {
		l.rear = nil
	}
	fmt.Println(e)
}

func (l *Qlinkedlist) first() {
	if l.isempty() {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println(l.frount.element)
}

func (l *Qlinkedlist) last() {
	if l.isempty() {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println(l.rear.element)
}
func (l *Qlinkedlist) display() {
	p := l.frount
	for p != nil {
		fmt.Print(p.element, " <--")
		p = p.next
	}
	fmt.Println()
}

func main() {
	Q := Qlinkedlist{}
	Q.enqueue(40)
	Q.enqueue(30)
	Q.enqueue(20)
	Q.enqueue(10)
	Q.display()
	fmt.Println("----------------------------------")
	Q.dqueue()
	Q.display()
	fmt.Println("----------------------------------")
	Q.first()
	Q.display()
	fmt.Println("----------------------------------")
	Q.last()
	Q.display()
	fmt.Println("----------------------------------")
}
