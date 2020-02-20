package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func parser(path string) header {
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
	h := header{}

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
			} else {
			}
			if i == 1 {
				tmp, err := strconv.Atoi(value)
				if err != nil {
					fmt.Println("ERROR\n")
				}
				h.scores = append(h.scores, tmp)
			}
		}
		i++
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	return h
}