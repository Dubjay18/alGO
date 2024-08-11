package stacks

import "fmt"

// Stack represents the stack structure(array).
type StackArr struct {
	Items []int // Items is a slice to store the stack elements.
}

// Push adds a new element to the top of the stack.
func (s *StackArr) Push(v int) {
	s.Items = append(s.Items, v) // Append the new element to the stack.
}

// Pop removes and returns the top element from the stack.
func (s *StackArr) Pop() int {
	if len(s.Items) == 0 { // If the stack is empty, return 0.
		return 0
	}
	// Remove the top element from the stack.
	v := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return v
}

// Peek returns the top element from the stack without removing it.
func (s *StackArr) Peek() int {
	if len(s.Items) == 0 { // If the stack is empty, return 0.
		return 0
	}
	return s.Items[len(s.Items)-1] // Return the top element of the stack.
}

// IsEmpty checks if the stack is empty.
func (s *StackArr) IsEmpty() bool {
	return len(s.Items) == 0 // Returns true if the stack is empty.
}

// Size returns the number of elements in the stack.
func (s *StackArr) Size() int {
	return len(s.Items) // Return the number of elements in the stack.
}

// Print prints the elements of the stack.
func (s *StackArr) Print() {
	ltr := ""                                // Initialize an empty string to store the stack values.
	for i := len(s.Items) - 1; i >= 0; i-- { // Traverse the stack from top to bottom.
		ltr += " " + fmt.Sprintf("%v -/n", s.Items[i]) // Append the value to the stack string.
	}
	println(ltr) // Print the stack string.
}

// TestStack is a function to test the stack implementation.
