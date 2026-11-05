package main

import (
	"fmt"
	"sync"
)

func node(i int, ch [][]chan float64, rounds int, x *float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for range rounds {
		for j := range ch {
			if ch[i][j] != nil {
				ch[i][j] <- *x
			}
		}
		sum, d := *x, 0
		for j := range ch {
			if ch[j][i] != nil {
				sum += <-ch[j][i]
				d++
			}
		}
		*x = sum / float64(d+1)
	}
}

func AverageFlooding(votes []float64, adj [][]int, rounds int) []float64 {
	n := len(votes)
	ch := make([][]chan float64, n)
	for i := range n {
		ch[i] = make([]chan float64, n)
		for _, j := range adj[i] {
			ch[i][j] = make(chan float64, 1)
		}
	}
	x := append([]float64{}, votes...)
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go node(i, ch, rounds, &x[i], &wg)
	}
	wg.Wait()
	return x
}

func main() {
	votes := []float64{1, 0, 1, 1, 0, 0, 1}
	n := len(votes)
	adj := make([][]int, n)
	for i := range n {
		adj[i] = []int{(i + n - 1) % n, (i + 1) % n}
	}
	x := AverageFlooding(votes, adj, 200)
	fmt.Println("Stari finale:", x)
	if x[0] > 0.5 {
		fmt.Println("rezultat vot: DA")
	} else {
		fmt.Println("Rezultat vot: NU")
	}
}
