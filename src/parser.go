package main

import (
	"strings"
	"strconv"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)
/*
func parser(path string) (header, library) {
	fptr := flag.String("fpath", path, "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	var i int = 0
	var lib int = -1
	var tmp int = 0
	h := header{}
	l := library{}

	for s.Scan() {
		str := strings.Split(s.Text(), " ")
		for index, value := range str {
			index = index
			if i == 0 {
				if index == 0 {
					h.books, err = strconv.Atoi(value)
				}
				if index == 1 {
					h.libraries, err = strconv.Atoi(value)
				}
				if index == 2 {
					h.days_scanning, err = strconv.Atoi(value)
				}
				if err != nil {
				}
			} else if i == 1 {
				tmp, err := strconv.Atoi(value)
				if err != nil {
					fmt.Println("ERROR\n")
				}
				h.scores = append(h.scores, tmp)
			} else {
				if i % 2 == 0 {
					fmt.Println(str)
					tmp = 1
				} else {
					fmt.Println(str)
				}
			}
		}
		if tmp == 1 {
			lib++
		}
		tmp = 0
		i++
		fmt.Println(lib)
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return h, l
}
*/

func parser(filename string) (header, []library) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
	}
	buf := bytes.NewBuffer(file)
	var i int = 0
	var lib int = 0
	var tmp int = 0
	var tmp_check = 0
	h := header{}
	l := library{}

	array_lib := []library{}

	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return h, array_lib
			}
		}
		if err != nil && err != io.EOF {
			return h, array_lib
		}
		str := strings.Split(line, " ")
		if tmp_check  == 1 {
			tmp_check = 0
		}
		for index, value := range str {
			index = index
			if i == 0 {
				if index == 0 {
					h.books, err = strconv.Atoi(value)
				}
				if index == 1 {
					h.libraries, err = strconv.Atoi(value)
				}
				if index == 2 {
					h.days_scanning, err = strconv.Atoi(value)
				}
				if err != nil {
				}
			} else if i == 1 {
				tmp, err := strconv.Atoi(value)
				if err != nil {
					//fmt.Println("ERROR\n")
				}
				h.scores = append(h.scores, tmp)
			} else {
				if i % 2 == 0 {
					l.name = lib
					if index == 0 {
 						l.books, err = strconv.Atoi(value)
					}
					if index == 1 {
						l.signup, err = strconv.Atoi(value)
					}
					if index == 2 {
						if strings.ContainsAny(value, "\n") {
							value = strings.Replace(value, "\n", "", len(value))
						}
						l.shipping, err = strconv.Atoi(value)
						//fmt.Println(l, index)
						if err == nil {
						}
					}
					tmp = 1
				} else {
					tmp, err := strconv.Atoi(value)
					if err != nil {
					}
					l.book_scores = append(l.book_scores, tmp)
					tmp_check = 1
				}
			}
		}
		if (tmp_check == 1) {
			array_lib = append(array_lib, l)
		}
		if tmp == 1 {
			lib++
		}
		tmp = 0
		i++
	}
	return h, array_lib
}

func main() {
	// a flat file that has 339276 lines of text in it for a size of 62.1 MB
	filename := "../test.txt"
	fmt.Println(parser(filename))
}