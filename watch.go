/*
*
*
*	count time and display it concurrently to test channels with goroutines.
*
*
*/


package main

import(
	"fmt"
	"time"
)

func main() {

	cMinute 	:= make(chan string)
	cSeconds 	:= make(chan string)
	cSecondsAlt := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			cSeconds <- "One second has passed."
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 15)
			cSecondsAlt <- "Fifteen seconds have passed."
		}
	}()
	go func() {
		for {
			time.Sleep(time.Minute)
			cMinute <- "One minute has passed."
		}
	}()

	for {
		select {
		case second := <- cSeconds:
			fmt.Println(second)
		case fifteen := <- cSecondsAlt:
			fmt.Println(fifteen)
		case minute := <- cMinute:
			fmt.Println(minute)
		}
	}
}