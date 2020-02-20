/*
*
*
*	Bubble sort + test on goroutines and wait-groups
*
*
*/

package main

// Bubble sort an array of integers.
func bubble_sort_libs(array []library) {

	for !is_sorted(array) {
		for i, j := 0, 1; j < len(array); i, j = i + 1, j + 1 {
			if array[i].shipping > array[j].shipping {
				swap(&array[i].shipping, &array[j].shipping)
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
func is_sorted(array []library) bool {

	for i := 0; i < len(array); i++ {
		if i + 1 < len(array) && array[i].shipping > array[i + 1].shipping {
			return false
		}
	}

	return true
}