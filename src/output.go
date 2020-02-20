package main

import (
	"fmt"
	"os"
	"strconv"
)

type output map[int][]int

// func main() {
// 	var out = map[int][]int{
// 		1: []int{1, 2, 3},
// 		2: []int{1, 1, 4},
// 	}
// 	createOutput(out)
// }

func createOutput(out output) {
	var nbrLib = len(out)
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Print("Error: cannot create file")
		return
	}
	file.WriteString(fmt.Sprintf("%d\n", nbrLib))
	for name, books := range out {
		var str = strconv.Itoa(name)
		str += " "
		str += strconv.Itoa(len(books))
		str += "\n"
		for b := range books {
			str += strconv.Itoa(b)
			str += " "
		}
		str += "\n"
		file.WriteString(str)
	}
	file.Close()
}
