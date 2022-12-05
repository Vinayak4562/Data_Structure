package main

import "fmt"

type node struct {
	element int
	next    *node
	prev    *node
}
type DoublyCirLL struct {
	head *node
	tail *node
	size int
}

func (l *DoublyCirLL) len() int {
	return l.size
}
func (l *DoublyCirLL) isempty() bool {
	return l.size == 0
}
func (l *DoublyCirLL) addfirst(e int) {
	newest := &node{e, nil, nil}
	if l.isempty() {
		l.head = newest
		l.tail = newest
	}
	newest.next = l.head
	l.head.prev = newest
	l.tail.next = newest
	newest.prev = l.tail
	l.head = newest
	l.size++
}
func (l *DoublyCirLL) addlast(e int) {
	newest := &node{e, nil, nil}
	if l.isempty() {
		l.head = newest
		l.tail = newest
	}
	p := l.head
	i := 0
	for i < l.len() {
		p = p.next
		i++
	}
	l.tail.next = newest
	newest.prev = l.tail
	l.tail = newest
	l.tail.next = l.head
	l.size++
}

func (l *DoublyCirLL) addany(e, position int) {
	newest := &node{e, nil, nil}
	if l.isempty() {
		l.head = newest
		l.tail = newest
	}
	if position == 1 {
		l.addfirst(e)
	} else if position > l.len() {
		l.addlast(e)
	} else {
		p := l.head
		i := 1
		for i < position-1 {
			p = p.next
			i++
		}
		newest.next = p.next
		p.next.prev = newest
		p.next = newest
		newest.prev = p
		l.size++
	}

}
func (l *DoublyCirLL) remvfirst() {
	if l.isempty() {
		fmt.Println("List is empty in remove first")
		return
	}
	e := l.head.element
	l.tail.next = l.head.next
	l.head = l.head.next
	l.head.prev = l.tail
	l.size--
	fmt.Println(e)
}
func (l *DoublyCirLL) removelast() {
	if l.isempty() {
		fmt.Println("List is empty in REMVLast")
		return
	}
	p := l.head
	i := 1
	for i < l.len()-1 {
		p = p.next
		i++
	}
	e := p.next.element
	p.next = l.head
	l.head.prev = p
	l.tail = p
	l.size--
	fmt.Println(e)
}
func (l *DoublyCirLL) removeany(position int) {
	if l.isempty() {
		fmt.Println("List is empty in REMVAny")
		return
	} else if position == 1 {
		l.remvfirst()
	} else if position >= l.len() {
		l.removelast()
	} else {
		p := l.head
		i := 1
		for i < position-1 {
			p = p.next
			i++
		}
		e := p.next.element
		p.next = p.next.next
		p.next.next.prev = p
		l.size--
		fmt.Println(e)
	}
}
func (l *DoublyCirLL) display() {
	p := l.head
	i := 0
	for i < l.len() {
		fmt.Print(p.element, "-->")
		p = p.next
		i++
	}
	fmt.Println()
}
func (l *DoublyCirLL) displyprev() {
	p := l.tail
	i := 0
	for i < l.len() {
		fmt.Print(p.element, "<--")
		p = p.prev
		i++
	}
	fmt.Println()
}

func main() {
	a := &DoublyCirLL{}
	a.addfirst(10)
	a.addfirst(12)
	a.addfirst(13)
	a.addfirst(14)
	a.display()
	a.displyprev()
	fmt.Println("---------------------------------------------")
	a.addlast(1)
	a.addlast(2)
	a.addlast(3)
	a.addlast(4)
	a.display()
	a.displyprev()
	fmt.Println("---------------------------------------------")
	a.addany(555, 2)
	a.display()
	a.displyprev()
	fmt.Println("---------------------------------------------")
	a.addany(666, 10)
	a.display()
	a.displyprev()
	fmt.Println("---------------------------------------------")
	a.remvfirst()
	a.display()
	a.displyprev()
	fmt.Println("---------------------------------------------")
	a.removelast()
	a.display()
	a.displyprev()
	fmt.Println("---------------------------------------------")
	a.removeany(2)
	a.display()
	a.displyprev()
	fmt.Println("---------------------------------------------")
	a.removeany(3)
	a.display()
	a.displyprev()
	fmt.Println("---------------------------------------------")
	a.removeany(6)
	a.display()
	a.displyprev()
	fmt.Println("---------------------------------------------")

}
