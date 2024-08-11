package test

import (
	"alGO/linkedlist"
	"alGO/queues"
	"alGO/stacks"
)

func TestLinkedList() {
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
func TestStack() {
	// Create a new stack.
	s := stacks.StackArr{}
	// Push elements to the stack.
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	// Print the stack elements.
	s.Print()
	// Pop an element from the stack.
	s.Pop()
	// Print the stack elements.
	s.Print()
	// Peek the top element of the stack.
	println("Top element:", s.Peek())
	// Check if the stack is empty.
	println("Is stack empty:", s.IsEmpty())
	// Print the number of elements in the stack.
	println("Stack size:", s.Size())
	sl := stacks.StackLL{}
	// Push nodes to the stack.
	sl.Push(10)
	sl.Push(20)
	sl.Push(30)
	sl.Push(40)
	// Print the stack nodes.
	sl.Print()
	// Pop a node from the stack.
	sl.Pop()
	// Print the stack nodes.
	sl.Print()
	// Peek the top node of the stack.
	println("Top node:", sl.Peek())
	// Check if the stack is empty.
	println("Is stack empty:", sl.IsEmpty())
	// Print the number of nodes in the stack.
	println("Stack size:", sl.Size())
}

func TestQueue() {
	// Create a new queue.
	q := queues.Queue{}
	// Enqueue elements to the queue.
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	// Print the queue elements.
	q.Print()
	// Dequeue an element from the queue.
	q.Dequeue()
	// Print the queue elements.
	q.Print()
	// Peek the front element of the queue.
	println("Front element:", q.Peek())
	// Check if the queue is empty.
	println("Is queue empty:", q.IsEmpty())
	// Print the number of elements in the queue.
	println("Queue size:", q.Size())

}
