package main

import(
	"os"
	"fmt"
	"io/ioutil"
)

func main() {

	if len(os.Args) < 2 {
		return
	}

	fmt.Println(len(os.Args) - 1)

	// Print all args
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("PASSED")
		func() {
			open_file(os.Args[i])
		}()
	}
}

// Open a file from path.
func open_file(filepath string) {

	data, error := ioutil.ReadFile(filepath)

	// Check if an error occured.
	if error != nil {
		fmt.Println("Error.")
		return
	}

	fmt.Println(data)
}