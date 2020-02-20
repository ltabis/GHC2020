package main

import(
	"fmt"
	// "errors"
)

func main() {
	array := []int{1, 2, 3, 1, 0, 8, 1}

	fmt.Println("before sort: ", array)
	bubble_sort(array)
	fmt.Println("after sort: ", array)
}

// Bubble sort an array of integers.
func bubble_sort(array []int) {

	for !is_sorted(array) {
		for i, j := 0, 1; j < len(array); i, j = i + 1, j + 1 {
			if array[i] > array[j] {
				swap(&array[i], &array[j])
			}
		}
	}
}

// Swaping two integers
func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

// Checks if an array is sorted
// Returns true if the array is sorted, false otherwise
func is_sorted(array []int) bool {

	for i := 0; i < len(array); i++ {
		if i + 1 < len(array) && array[i] > array[i + 1] {
			return false
		}
	}

	return true
}