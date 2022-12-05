package main

import "fmt"

type node struct {
	data int
	next *node
}

type linked struct {
	head *node
	tail *node
}

func (l *linked) addlast(n int) {
	new := &node{n, nil}
	if l.head == nil {
		l.head = new
	} else {
		l.tail.next = new
	}
	l.tail = new

}

func (l *linked) disply() {
	disp := l.head
	for disp != nil {
		fmt.Println(disp.data)
		disp = disp.next

	}

}

func main() {
	appd := linked{}
	appd.addlast(10)
	appd.addlast(12)
	appd.addlast(13)
	appd.addlast(14)

	appd.disply()

}
