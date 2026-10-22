package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}
	mapSlice := map[int][]string{}
	for k, v := range myMap {
		_, ok := mapSlice[v]
		if !ok {
			mapSlice[v] = []string{k}
		} else {
			mapSlice[v] = append(mapSlice[v], k)
		}
	}
	fmt.Println(mapSlice)
}
