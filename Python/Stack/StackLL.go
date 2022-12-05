package main

import "fmt"

type node struct {
	element int
	next    *node
}
type StackLL struct {
	first *node
	rear  *node
	size  int
}

func (l *StackLL) len() int {
	return l.size
}

func (l *StackLL) isempty() bool {
	return l.size == 0
}

func (l *StackLL) addfrnt(e int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.first = newest
		l.rear = newest
	}
	newest.next = l.first
	l.first = newest
	l.size++
}
func (l *StackLL) remvove() {
	if l.isempty() {
		fmt.Println("Stack is empty")
		return
	}
	e := l.first.element
	l.first = l.first.next
	l.size--
	fmt.Println(e)
}

func (l *StackLL) display() {
	p := l.first
	i := 0
	for i < l.len() {
		fmt.Print(p.element, "-->")
		p = p.next
		i++

	}
	fmt.Println()
}
func main() {
	s := &StackLL{}
	s.addfrnt(10)
	s.addfrnt(20)
	s.addfrnt(30)
	s.addfrnt(40)
	s.display()
	s.remvove()
	s.display()
}
