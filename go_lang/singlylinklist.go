package main

import "fmt"

type node struct {
	data int
	next *node
}
type linked struct {
	head *node
	tail *node
	size int
}

func (l *linked) len() int {
	return l.size

}

func (l *linked) isempty() bool {
	return l.size == 0
}

func (l *linked) addfirst(e int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.head = newest
		l.tail = newest
	} else {
		l.tail.next = newest
	}
	l.tail = newest
	l.size += 1

}

func (l *linked) disply() {
	p := l.head
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}

}

func main() {
	n := linked{}

	n.addfirst(2)
	n.addfirst(3)
	n.addfirst(4)
	n.addfirst(5)
	n.disply()

}
