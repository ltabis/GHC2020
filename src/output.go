package main

import (
	"fmt"
	"os"
	"strconv"
)

type output map[int][]int

func createOutput(out output) {
	var nbrLib = len(out)
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Print("Error: cannot create file")
		return
	}
	_, err = file.WriteString(fmt.Sprintf("%d\n", nbrLib))
	if err != nil {
		fmt.Print("Error: cannot write")
	}
	for name, books := range out {
		var str = strconv.Itoa(name)
		str += strconv.Itoa(len(books))
		str += "\n"
		for b := range books {
			str += strconv.Itoa(b)
		}
		str += "\n"
		file.WriteString(str)
	}
	file.Close()
}
