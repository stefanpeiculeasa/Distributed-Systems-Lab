package main

import (
	"fmt"
	"sync"
)

func node(i int, id int, ch [][]chan int, rounds int, maxId *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*maxId = id
	for range rounds {
		for j := range ch {
			if ch[i][j] != nil {
				ch[i][j] <- *maxId
			}
		}
		for j := range ch {
			if ch[j][i] != nil {
				*maxId = max(*maxId, <-ch[j][i])
			}
		}
	}
}

func GraphFlooding(ids []int, adj [][]int) int {
	n := len(ids)
	ch := make([][]chan int, n)
	for i := range n {
		ch[i] = make([]chan int, n)
		for _, j := range adj[i] {
			ch[i][j] = make(chan int, 1)
		}
	}
	maxId := make([]int, n)
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go node(i, ids[i], ch, n-1, &maxId[i], &wg)
	}
	wg.Wait()
	leader := -1
	for i := range n {
		fmt.Printf("nod %d (id %d): max_id = %d\n", i, ids[i], maxId[i])
		if maxId[i] == ids[i] {
			leader = ids[i]
		}
	}
	return leader
}

func main() {
	ids := []int{4, 19, 7, 2, 11}
	adj := [][]int{{1}, {2, 3}, {0}, {4}, {0, 2}}
	fmt.Println("Lider:", GraphFlooding(ids, adj))
}
