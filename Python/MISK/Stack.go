package main

import "fmt"

type node struct {
	element int
	next    *node
}
type Stack struct {
	first *node
	rear  *node
	size  int
}

func (l *Stack) len() int {
	return l.size
}
func (l *Stack) isempty() bool {
	return l.size == 0
}
func (l *Stack) add(e int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.first = newest
		l.rear = newest
	}
	l.rear.next = newest
	l.rear = newest
	l.size++
}

func (l *Stack) display() {
	p := l.first
	i := 0
	for p != nil {
		fmt.Print(p.element, "-->")
		p = p.next
		i++
	}
	fmt.Print()
}
func main() {
	s := &Stack{}
	s.add(10)
	s.add(20)
	s.add(30)
	s.display()
}
