package stacks

import (
	"alGO/linkedlist"
	"fmt"
)

// Stack represents the stack structure(linkedlists).

type StackLL struct {
	Head   *linkedlist.Node // Head is a pointer to the first node of the stack.
	length int              // length keeps track of the number of nodes in the stack.
	Tail   *linkedlist.Node // Tail is a pointer to the last node of the stack.
}

// CheckIFEmpty checks if the stack is empty.
func CheckIFEmpty(l *StackLL) bool {
	return l.Head == nil // Returns true if the Head is nil, indicating the stack is empty.
}

// print prints the values of all nodes in the stack.
func (l *StackLL) Print() {
	if CheckIFEmpty(l) { // If the stack is empty, return without printing anything.
		return
	}
	saveNode := l.Head    // Start with the Head node.
	lstr := ""            // Initialize an empty string to store the stack values.
	for saveNode != nil { // Traverse the stack until the end.
		lstr += " " + fmt.Sprintf("%v -/n", saveNode.Value) // Append the value to the stack string.
		saveNode = saveNode.Next                            // Move to the next node.
	}
	println(lstr + " nil") // Print the stack string.
}

// CheckLength calculates and returns the length of the stack.
func (l *StackLL) CheckLength() int {
	if CheckIFEmpty(l) { // If the stack is empty, return 0.
		return 0
	}
	saveNode := l.Head    // Start with the Head node.
	for saveNode != nil { // Traverse the stack to count the nodes.
		l.length++               // Increment the length counter for each node.
		saveNode = saveNode.Next // Move to the next node.
	}
	return l.length // Return the total number of nodes.
}

// Push adds a new node with value v at the top of the stack.
func (l *StackLL) Push(v int) {
	if CheckIFEmpty(l) { // If the stack is empty, create a new Head node.
		l.Head = &linkedlist.Node{
			Value: v,
			Next:  nil,
		}
		l.Tail = l.Head // The Head node is also the Tail node.
		return
	}
	// Push a new node at the top of the stack.
	l.Head = &linkedlist.Node{
		Value: v,
		Next:  l.Head,
	}
}

// Pop removes and returns the top node from the stack.
func (l *StackLL) Pop() int {
	if CheckIFEmpty(l) { // If the stack is empty, return 0.
		return 0
	}
	// Remove the top node from the stack.
	v := l.Head.Value
	l.Head = l.Head.Next
	return v
}

// Peek returns the top node from the stack without removing it.
func (l *StackLL) Peek() int {
	if CheckIFEmpty(l) { // If the stack is empty, return 0.
		return 0
	}
	return l.Head.Value // Return the value of the top node.
}

// IsEmpty checks if the stack is empty.
func (l *StackLL) IsEmpty() bool {
	return l.Head == nil // Returns true if the stack is empty.
}

// Size returns the number of nodes in the stack.
func (l *StackLL) Size() int {
	return l.CheckLength() // Return the number of nodes in the stack.
}
