package queues

import "fmt"

type Queue struct {
	Items []int // Items is a slice to store the queue elements.
}

// Enqueue adds a new element to the end of the queue.
func (q *Queue) Enqueue(v int) {
	q.Items = append(q.Items, v) // Append the new element to the queue.
}

// Dequeue removes and returns the front element from the queue.
func (q *Queue) Dequeue() int {
	if len(q.Items) == 0 { // If the queue is empty, return 0.
		return 0
	}
	// Remove the front element from the queue.
	v := q.Items[0]
	q.Items = q.Items[1:]
	return v
}

// Peek returns the front element from the queue without removing it.
func (q *Queue) Peek() int {
	if len(q.Items) == 0 { // If the queue is empty, return 0.
		return 0
	}
	return q.Items[0] // Return the front element of the queue.
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0 // Returns true if the queue is empty.
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.Items) // Return the number of elements in the queue.
}

// Print prints the elements of the queue.
func (q *Queue) Print() {
	ltr := ""                           // Initialize an empty string to store the queue values.
	for i := 0; i < len(q.Items); i++ { // Traverse the queue from front to end.
		ltr += " " + fmt.Sprintf("%v -/n", q.Items[i]) // Append the value to the queue string.
	}
	println(ltr) // Print the queue string.
}
