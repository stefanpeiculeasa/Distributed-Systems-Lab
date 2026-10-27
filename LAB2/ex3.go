package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "mesaj1"
	}()
	go func() {
		ch <- "mesaj2"
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
