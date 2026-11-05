package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool
}

func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- Msg{id, false}
	for {
		m := <-in
		if m.Done {
			if !*isLeader {
				out <- m
			}
			return
		}
		if m.Value > id {
			out <- m
		} else if m.Value == id {
			*isLeader = true
			out <- Msg{id, true}
		}
	}
}

func RingFlooding(ids []int) int {
	n := len(ids)
	ch := make([]chan Msg, n)
	for i := range ch {
		ch[i] = make(chan Msg, n+1)
	}
	isLeader := make([]bool, n)
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go node(ids[i], ch[i], ch[(i+1)%n], &isLeader[i], &wg)
	}
	wg.Wait()
	for i := range n {
		if isLeader[i] {
			return ids[i]
		}
	}
	return -1
}

func main() {
	fmt.Println("Lider:", RingFlooding([]int{3, 17, 5, 42, 8, 1}))
}
