package output

// package output

import (
	"fmt"
	"os"
)

type output map[string][]int

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
		file.WriteString(fmt.Sprintf("%s %d\n", name, len(books)))
		for b := range books {
			file.WriteString(fmt.Sprintf("%d ", b))
		}
		file.WriteString("\n")
	}
	file.Close()
}
