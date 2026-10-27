package main

import (
	"fmt"
	"time"
)

type pkt struct {
	x int
	y int
}

func robot1(traseu []pkt, ch1 chan<- pkt, ch2 <-chan pkt, ch3 chan string) {
	other := []pkt{}
	for _, val := range traseu {
		ch1 <- val
		time.Sleep(1 * time.Second)
	}
	close(ch1)
	for elem := range ch2 {
		other = append(other, elem)
		fmt.Printf("Received from r2: x = %d, y = %d\n", elem.x, elem.y)
	}
	ch3 <- "robot1 done"
}

func robot2(traseu []pkt, ch1 <-chan pkt, ch2 chan<- pkt, ch3 chan string) {
	other := []pkt{}
	for _, val := range traseu {
		ch2 <- val
		time.Sleep(1 * time.Second)
	}
	close(ch2)
	for elem := range ch1 {
		other = append(other, elem)
		fmt.Printf("received from r1: x = %d, y = %d\n", elem.x, elem.y)
	}
	ch3 <- "robot2 done"
}

func main() {
	traseu := []pkt{{x: 1, y: 2}, {x: 2, y: 3}, {x: 4, y: 3}, {x: 5, y: 5}, {x: 7, y: 6}, {x: 8, y: 8}}
	ch1 := make(chan pkt, 100)
	ch2 := make(chan pkt, 100)
	ch3 := make(chan string)
	go robot1(traseu, ch1, ch2, ch3)
	go robot2(traseu, ch1, ch2, ch3)
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)
}
