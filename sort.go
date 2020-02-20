/*
*
*
*	Bubble sort + test on goroutines and wait-groups
*
*
*/

package main

import(
	"fmt"
	"sync"
	// "errors"
)

func main() {

	// Creating a wait group for parallelism.
	var wg sync.WaitGroup

	// We're gonna have only one goroutine to wait before exiting the program.
	wg.Add(1)

	array := []int{1, 2, 3, 1, 0, 8, 1}
	array2 := []int{1, 2, 3, 1, 0, 8, 1}

	fmt.Println("before sort 1: ", array)
	fmt.Println("before sort 2: ", array2)

	// goroutine + anonymous function (lambda).
	go func() {
		bubble_sort(array)

		// Goroutine as finished, main thread can close.
		wg.Done()
	}()
	bubble_sort(array2)

	// Waiting for the goroutine to finish.
	wg.Wait()

	fmt.Println("after sort 1: ", array)
	fmt.Println("after sort 2: ", array2)
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