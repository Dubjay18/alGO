package linkedlist

import "fmt"

type DNode struct {
	Value int
	Next  *DNode
	Prev  *DNode
}

type DlinkedList struct {
	Head   *DNode
	Tail   *DNode
	length int
}

func (l *DlinkedList) CheckLength() int {
	if l.CheckIFEmpty() { // If the list is empty, return 0.
		return 0
	}
	saveNode := l.Head    // Start with the Head node.
	for saveNode != nil { // Traverse the list to count the nodes.
		l.length++               // Increment the length counter for each node.
		saveNode = saveNode.Next // Move to the next node.
	}
	return l.length // Return the total number of nodes.
}

func (d *DlinkedList) CheckIFEmpty() bool {
	return d.Head == nil
}

func (d *DlinkedList) Insert(v int, idx int) {
	if d.CheckLength() < idx { // If the index is out of range, print an error message.
		println("Index out of range")
		return
	}
	saveNode := d.Head         // Start with the Head node.
	for j := 0; j < idx; j++ { // Traverse the list to the node just before the insertion point.
		saveNode = saveNode.Next
	}
	// Insert the new node at the specified index.
	saveNode.Next = &DNode{
		Value: v,
		Next:  saveNode.Next, // The new node points to the next node in the list.
		Prev:  saveNode,
	}
}

func (d *DlinkedList) Prepend(v int) {
	if d.CheckIFEmpty() { // If the list is empty, create a new Head node.
		d.Head = &DNode{
			Value: v,
			Next:  nil,
			Prev:  nil,
		}
		d.Tail = d.Head // The Head node is also the Tail node.
		return
	}
	// Prepend a new node at the beginning of the list.
	newNode := &DNode{
		Value: v,
		Next:  nil, // The new node points to the old Head node.
		Prev:  nil,
	}

	newNode.Next = d.Head // The new node becomes the new Head.
	d.Head.Prev = newNode
	d.Head = newNode
}

func (d *DlinkedList) Append(v int) {
	if d.CheckIFEmpty() { // If the list is empty, create a new Head node.
		d.Head = &DNode{
			Value: v,
			Next:  nil,
			Prev:  nil,
		}
		d.Tail = d.Head // The Head node is also the Tail node.
		return
	}

	d.Tail.Next = &DNode{ // Create a new node and make it the new Head.
		Value: v,
		Next:  nil,
		Prev:  d.Tail,
	}
	d.Tail = d.Tail.Next
}

func (d *DlinkedList) Delete(idx int) {
	if d.CheckLength() < idx { // If the index is out of range, print an error message.
		println("Index out of range")
		return
	}
	saveNode := d.Head           // Start with the Head node.
	for j := 0; j < idx-1; j++ { // Traverse the list to the node just before the deletion point.
		saveNode = saveNode.Next
	}
	// Delete the node at the specified index.
	saveNode.Next = saveNode.Next.Next
}

func (d *DlinkedList) Print() {
	if d.CheckIFEmpty() { // If the list is empty, print an error message.
		println("List is empty")
		return
	}
	saveNode := d.Head    // Start with the Head node.
	lstr := ""            // Initialize an empty string to store the list values.
	for saveNode != nil { // Traverse the list to print the nodes.
		lstr += " " + fmt.Sprintf("%v -><-", saveNode.Value)
		saveNode = saveNode.Next // Move to the next node.
	}
	println(lstr + " nil") // Print the list string.
}

func (d *DlinkedList) Search(v int) {
	if d.CheckIFEmpty() { // If the list is empty, print an error message.
		println("List is empty")
		return
	}
	saveNode := d.Head                 // Start with the Head node.
	for i := 0; saveNode != nil; i++ { // Traverse the list to search for the value.
		if saveNode.Value == v { // If the value is found, print the index.
			println("Value found at index:", i)
			return
		}
		saveNode = saveNode.Next // Move to the next node.
	}
	println("Value not found") // If the value is not found, print an error message.
}

func (d *DlinkedList) Reverse() {
	if d.CheckIFEmpty() { // If the list is empty, print an error message.
		println("List is empty")
		return
	}
	saveNode := d.Head    // Start with the Head node.
	for saveNode != nil { // Traverse the list to the end.
		saveNode = saveNode.Next // Move to the next node.
	}
	// Traverse the list in reverse order to print the nodes.
	for saveNode != nil {
		println(saveNode.Value)  // Print the value of the current node.
		saveNode = saveNode.Prev // Move to the previous node.
	}
}
