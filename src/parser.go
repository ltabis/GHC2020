package main

import (
	"strings"
	"strconv"
	"bytes"
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
func parser(filename string) (header, library) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
	}
	buf := bytes.NewBuffer(file)
	var i int = 0
	var lib int = -1
	var tmp int = 0
	h := header{}
	l := library{}

	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return h, l
			}
		}
		if err != nil && err != io.EOF {
			return h, l
		}
		str := strings.Split(line, " ")
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
//					fmt.Println("ERROR\n")
				}
				h.scores = append(h.scores, tmp)
			} else {
				l = l
				if i % 2 == 0 {
//					fmt.Println(str)
					tmp = 1
				} else {
//					fmt.Println(str)
				}
			}
		}
		if tmp == 1 {
			lib++
		}
		tmp = 0
		i++
	}
	return h, l
}