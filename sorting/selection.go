package sorting

// Selection sorts the given array of integers in ascending order using the selection sort algorithm.
// It takes an array as input and returns the sorted array.
func Selection(arr []int) []int {
	// Get the length of the array
	n := len(arr)

	// Iterate through the array
	for i := 0; i < n-1; i++ {
		// Find the minimum element in the unsorted part of the array
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		// Swap the minimum element with the first element of the unsorted part
		arr[i], arr[min] = arr[min], arr[i]
	}

	// Return the sorted array
	return arr
}
