package main

type node struct {
	element int
	next    *node
	prev    *node
}

type DublyLinkedList struct {
	head *node
	tail *node
	size int
}

func (l *DublyLinkedList) len() int {
	return l.size
}
