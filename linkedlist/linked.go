package linkedlist

import "fmt"

// Node represents an individual element in the linked list.
type Node struct {
	Value int   // Value stores the data of the node.
	Next  *Node // Next is a pointer to the next node in the list.
}

// SLinkedList represents the Singley linked list structure.
type SLinkedList struct {
	Head   *Node // Head is a pointer to the first node of the list.
	length int   // length keeps track of the number of nodes in the list.
}

// CheckIFEmpty checks if the linked list is empty.
func CheckIFEmpty(l *SLinkedList) bool {
	return l.Head == nil // Returns true if the Head is nil, indicating the list is empty.
}

// print prints the values of all nodes in the linked list.
func (l *SLinkedList) Print() {
	if CheckIFEmpty(l) { // If the list is empty, return without printing anything.
		return
	}
	saveNode := l.Head    // Start with the Head node.
	lstr := ""            // Initialize an empty string to store the list values.
	for saveNode != nil { // Traverse the list until the end.
		lstr += " " + fmt.Sprintf("%v ->", saveNode.Value) // Append the value to the list string.
		saveNode = saveNode.Next                           // Move to the next node.
	}
	println(lstr + " nil") // Print the list string.
}

// CheckLength calculates and returns the length of the linked list.
func (l *SLinkedList) CheckLength() int {
	if CheckIFEmpty(l) { // If the list is empty, return 0.
		return 0
	}
	saveNode := l.Head    // Start with the Head node.
	for saveNode != nil { // Traverse the list to count the nodes.
		l.length++               // Increment the length counter for each node.
		saveNode = saveNode.Next // Move to the next node.
	}
	return l.length // Return the total number of nodes.
}

// prepend adds a new node with value v at the beginning of the linked list.
func (l *SLinkedList) Prepend(v int) {
	if CheckIFEmpty(l) { // If the list is empty, create a new Head node.
		l.Head = &Node{
			Value: v,
			Next:  nil,
		}
		return
	}
	// Prepend a new node at the beginning of the list.
	l.Head.Next = &Node{
		Value: v,
		Next:  nil,
	}
}

// append adds a new node with value v at the end of the linked list.
func (l *SLinkedList) Append(v int) {
	if CheckIFEmpty(l) { // If the list is empty, create a new Head node.
		l.Head = &Node{
			Value: v,
			Next:  nil,
		}
		return
	}
	saveNode := l.Head // Start with the Head node.
	l.Head = &Node{    // Create a new node and make it the new Head.
		Value: v,
		Next:  saveNode, // The old Head node becomes the next node.
	}
}

// insert adds a new node with value v at the specified index i in the linked list.
func (l *SLinkedList) Insert(v int, i int) {
	if l.CheckLength() < i { // If the index is out of range, print an error message.
		println("Index out of range")
		return
	}
	saveNode := l.Head       // Start with the Head node.
	for j := 0; j < i; j++ { // Traverse the list to the node just before the insertion point.
		saveNode = saveNode.Next
	}
	// Insert the new node at the specified index.
	saveNode.Next = &Node{
		Value: v,
		Next:  saveNode.Next, // The new node points to the next node in the list.
	}
}

// delete removes the node at the specified index i from the linked list.
func (l *SLinkedList) Delete(i int) {
	if l.CheckLength() < i { // If the index is out of range, print an error message.
		println("Index out of range")
		return
	}
	saveNode := l.Head         // Start with the Head node.
	for j := 0; j < i-1; j++ { // Traverse the list to the node just before the deletion point.
		saveNode = saveNode.Next
	}
	// Delete the node at the specified index.
	saveNode.Next = saveNode.Next.Next
}
