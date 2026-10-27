package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch2 <- <-ch1
	}()
	go func() {
		ch1 <- "hello"
	}()
	fmt.Println(<-ch2)

}
