package main

import "fmt"

type node struct {
	element int
	next    *node
}

type SinglyCir struct {
	head *node
	tail *node
	size int
}

func (l *SinglyCir) len() int {
	return l.size
}

func (l *SinglyCir) isempty() bool {
	return l.size == 0
}

func (l *SinglyCir) Addfirst(e int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.head = newest
		l.tail = newest
	}
	newest.next = l.head
	l.head = newest
	l.tail.next = l.head
	l.size++
}

func (l *SinglyCir) Addlast(e int) {
	newest := &node{e, nil}
	if l.isempty() {
		l.head = newest
		l.tail = newest
	}
	p := l.head
	i := 0
	for i < l.len()-1 {
		p = p.next
		i++
	}
	l.tail.next = newest
	newest.next = l.head
	l.tail = newest
	l.size++
}

func (l *SinglyCir) Addany(e, position int) {

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
func (l *SinglyCir) remvfirst() {
	if l.isempty() {
		fmt.Println("List is empty remvfirst")
		return
	}
	l.tail.next = l.head.next
	l.head = l.head.next
	l.size--

}
func (l *SinglyCir) remvlast() {
	if l.isempty() {
		fmt.Println("List is empty remvlast")
		return
	}
	p := l.head
	i := 1
	for i < l.len()-1 {
		p = p.next
		i++
	}
	e := p.next.element
	p = l.head
	l.tail = p
	l.size--
	fmt.Println(e)

}
func (l *SinglyCir) remvany(position int) {
	if l.isempty() {
		fmt.Println("List is empty")
		return
	}
	p := l.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	e := p.next.element
	p.next = p.next.next
	l.size--
	fmt.Println(e)
}

func (l *SinglyCir) display() {
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
	a := &SinglyCir{}
	a.Addfirst(10)
	a.Addfirst(20)
	a.Addfirst(30)
	a.Addfirst(40)
	a.display()
	fmt.Println("---------------------------------------------")
	a.Addlast(1)
	a.Addlast(2)
	a.Addlast(3)
	a.Addlast(4)
	a.display()
	fmt.Println("---------------------------------------------")
	a.remvany(2)
	a.display()
	// a.Addany(555, 2)
	// a.display()
	// fmt.Println("---------------------------------------------")
	// a.Addany(666, 3)
	// a.Addany(777, 4)
	// a.display()
	// fmt.Println("---------------------------------------------")
	// a.Addany(777, 15)
	// a.display()
	// fmt.Println("---------------------------------------------")
	// a.remvfirst()
	// a.display()
	// a.remvfirst()
	// a.display()
	// fmt.Println("---------------------------------------------")
	// a.remvlast()
	// a.display()
	// a.remvlast()
	// a.display()
	// fmt.Println("---------------------------------------------")
	// a.remvany(2)
	// a.display()
	// fmt.Println("---------------------------------------------")
	// a.remvany(2)
	// a.remvany(1)
	// a.remvany(1)
	// a.remvany(1)
	// a.remvany(1)
	// a.remvany(1)
	// a.remvany(1)
	// a.display()
	// fmt.Println("---------------------------------------------")
	// a.remvfirst()
	// a.display()
	// a.remvlast()
	// a.display()
	// a.Addany(2, 1)
	// a.display()
}
