package main

import "fmt"

type node struct {
	element int
	next    *node
}
type Queue struct {
	first *node
	rear  *node
	size  int
}

func (l *Queue) len() int {
	return l.size
}
func (l *Queue) isempty() bool {
	return l.size == 0
}
func (l *Queue) push(e int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.first = newest
		l.rear = newest
	}
	newest.next = l.first
	l.first = newest
	l.size++
}
func (l *Queue) pop() {
	if l.isempty() {
		fmt.Println("Queue is empty")
		return
	}
	e := l.first.element
	l.first = l.first.next
	l.size--
	fmt.Println(e)
}
func (l *Queue) top() int {
	if l.isempty() {
		fmt.Println("Queue is empty")

	}
	return l.first.element
}

func (l *Queue) display() {
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
	q := &Queue{}
	q.push(10)
	q.push(20)
	q.push(30)
	q.push(40)
	q.display()
	q.pop()
	// q.pop()
	// q.pop()
	// q.pop()
	// q.pop()
	q.display()

	fmt.Println(q.top())

}
