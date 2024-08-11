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
	l.Append(50)
	l.Append(60)
	l.Print()
	l.Delete(2)
	l.Print()
}
