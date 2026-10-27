package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 1000; i++ {
		}
		ch <- "done"
	}()
	fmt.Println(<-ch)

}
