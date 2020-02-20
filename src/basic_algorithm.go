package main

import(
	"fmt"
)

func main() {

	head, libs := parser("b_read_on.txt")

	fmt.Println(head)
	fmt.Println(libs)

	// out := output{
	// 	"rsc": {3711},
	// 	"r":   {2138},
	// 	"gri": {1908},
	// 	"adg": {912},
	// }

	// fmt.Println("out: ")
	// fmt.Println(out)

	// bubble_sort_libs(libs)
	// fmt.Println("libs: ")
	// fmt.Println(libs)

	// // createOutput(out)
	// fmt.Println("header: ")
	// fmt.Println(head)

	// run(&head, &libs)
}

func run(head *header, libs *[]library) output {

	var books_shiped []int
	current_waiting_days := 0

	for i := 0; i < head.days_scanning; i++ {
		for _, lib := range *libs {
			if lib.unlocked == true {
				append_books(&lib, &books_shiped)
			}
		}

		// Reserve the library
		current_waiting_days--
		if current_waiting_days < 0 {
			current_waiting_days = 0
		}
	}

	return nil
}

func append_books(lib *library, books_shiped *[]int) {

	var book_taken []int

	for b := 0; b < lib.shipping; b++ {

		for _, book := range lib.book_scores {

			if !in(book, *books_shiped) {
				
				lib.book_scores = remove_book(book, lib.book_scores)
			}
		}
	}

	for _, book := range book_taken {
		
		if !in(book, book_taken) {

			lib.book_scores = remove_book(book, lib.book_scores)
		}
	}
}

func in(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// Removes book from an array
func remove_book(i int, s []int) []int {
    s[i] = s[len(s) - 1]
    return s[:len(s) - 1]
}