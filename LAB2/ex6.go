package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "robot1"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "robot2"
	}()
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}

}
