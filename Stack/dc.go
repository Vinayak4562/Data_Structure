package main

import (
	"fmt"
)

type node struct {
	elememt int
	next    *node
	prev    *node
}

type DoubCir struct {
	head *node
	tail *node
	size int
}

func (l *DoubCir) len() int {
	return l.size
}

func (l *DoubCir) isempty() bool {
	return l.size == 0
}

func (l *DoubCir) addfirst(e int) {
	newest := &node{e, nil, nil}
	if l.isempty() {
		newest.prev = newest
		l.head = newest
		l.tail = newest
	} else {
		newest.next = l.head
		l.head.prev = newest
		l.head = newest
		l.tail.next = l.head
		l.head.prev = l.tail

	}
	l.size++
}

func (l *DoubCir) addlast(e int) {
	newest := &node{e, nil, nil}
	if l.isempty() {
		newest.next = newest
		newest.prev = newest
		l.head = newest
		l.tail = newest
	} else {
		l.tail.next = newest
		newest.prev = l.tail
		l.tail = newest
		l.tail.next = l.head
		l.head.prev = l.tail
	}
	l.size++

}

func (l *DoubCir) addany(e, position int) {
	newest := &node{e, nil, nil}
	p := l.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	newest.next = p.next
	p.next.prev = newest
	p.next = newest
	p.next.prev = p
	l.size++
}

func (l *DoubCir) remlast() {
	if l.isempty() {
		fmt.Println("List is empty")
		return

	}
	e := l.tail.elememt
	l.tail = l.tail.prev
	l.tail.next = l.head
	l.size--
	fmt.Println(e)
}

func (l *DoubCir) remany(position int) {
	p := l.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	e := p.next.elememt
	p.next = p.next.next
	p.next.prev = p
	l.size--
	fmt.Println(e)
}
func (l *DoubCir) remfirst() {
	if l.isempty() {
		fmt.Println("List is empty")
		return
	} else {
		e := l.head.elememt
		l.head = l.head.next
		// l.head.prev = l.tail
		l.tail.next = l.head

		l.size--
		fmt.Println(e)
	}
}

func (l *DoubCir) display() {
	p := l.head
	i := 0
	for i < l.len() {
		fmt.Print(p.elememt, "-->")
		p = p.next
		i += 1

	}
	fmt.Println()

}

func (l *DoubCir) displayRev() {
	p := l.tail
	i := 0
	for i < l.len() {
		fmt.Print(p.elememt, "<--")
		p = p.prev
		i += 1
	}
	fmt.Println()
}

func main() {
	S := &DoubCir{}
	fmt.Println("------------------------add first method----------------------")
	S.addfirst(10)
	S.addfirst(20)
	S.addfirst(30)
	S.display()
	S.displayRev()
	fmt.Println("------------------------add last method----------------------")
	S.addlast(1)
	S.addlast(2)
	S.display()
	S.displayRev()
	fmt.Println("------------------------add any method----------------------")
	S.addany(2, 3)
	S.display()
	S.displayRev()
	fmt.Println("------------------------aremove last method----------------------")
	S.remlast()
	S.display()
	S.displayRev()
	fmt.Println("------------------------aremove any method----------------------")
	S.remany(2)
	S.display()
	S.displayRev()
	fmt.Println("------------------------aremove first method----------------------")
	S.remfirst()
	S.display()
	S.displayRev()
}
