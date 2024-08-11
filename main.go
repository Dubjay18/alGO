package main

import (
	"alGO/linkedlist"
)

func main() {
	// linked list
	l := linkedlist.SLinkedList{}
	l.Prepend(10)
	l.Prepend(20)
	l.Prepend(30)
	l.Prepend(40)
	l.Append(80)
	l.Print()
	l.Delete(2)
	l.Print()
	l.Search(40)
	// doubly linked list
	d := linkedlist.DlinkedList{}
	d.Prepend(10)
	d.Prepend(20)
	d.Prepend(30)
	d.Prepend(40)
	d.Append(50)
	d.Append(60)
	d.Print()

}
