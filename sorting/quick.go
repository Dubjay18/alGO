package sorting




func QuickSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	left, right := 0, n-1
	pivot := n / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	return arr
}