package main

import "fmt"

type node struct {
	element int
	next    *node
}
type DEQ struct {
	frount *node
	rear   *node
	size   int
}

func (q *DEQ) len() int {
	return q.size
}
func (q *DEQ) isempty() bool {
	return q.size == 0
}
func (q *DEQ) addfirst(e int) {
	newest := &node{e, nil}
	if q.isempty() {
		q.frount = newest
		q.rear = newest
	}
	newest.next = q.frount
	q.frount = newest
	q.size++
}
func (q *DEQ) addlast(e int) {
	newest := &node{e, nil}
	if q.isempty() {
		q.frount = newest
		q.rear = newest
	}
	q.rear.next = newest
	q.rear = newest
	q.size++
}
func (q *DEQ) remvfirst() {
	if q.isempty() {
		fmt.Println("Queue is empty")
		return
	}
	e := q.frount.element
	q.frount = q.frount.next
	q.size--
	fmt.Println(e)
}
func (q *DEQ) remvlast() {
	if q.isempty() {
		fmt.Println("queue is empty")
		return
	}
	p := q.frount
	i := 1
	for i < q.len()-1 {
		p = p.next
		i++
	}
	e := p.next.element
	q.rear = p
	p.next = nil
	q.size--
	fmt.Println(e)
}
func (q *DEQ) top() int {
	if q.isempty() {
		fmt.Println("Quequ is empty")
	}
	return q.frount.element
}
func (q *DEQ) bottom() int {
	if q.isempty() {
		fmt.Println("Quequ is empty")
	}
	return q.rear.element
}
func (q *DEQ) display() {
	p := q.frount
	i := 0
	for i < q.len() {
		fmt.Print(p.element, "-->")
		p = p.next
		i++
	}
	fmt.Println()
}
func main() {
	d := &DEQ{}
	d.addfirst(10)
	d.addfirst(20)
	d.display()
	d.addlast(111)
	d.addlast(222)
	d.display()
	d.remvfirst()
	d.display()
	d.remvlast()
	d.display()
	fmt.Println(d.top())
	fmt.Println(d.bottom())
}
