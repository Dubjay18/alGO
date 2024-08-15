package sorting

import "fmt"

func Merged(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	fmt.Println("mid:", mid)
	left := Merged(arr[:mid])
	right := Merged(arr[mid:])
	fmt.Println("left:", left)
	fmt.Println("right:", right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	for i < len(left) {
		final = append(final, left[i])
		i++
	}
	for j < len(right) {
		final = append(final, right[j])
		j++
	}
	return final
}
