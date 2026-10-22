package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}
	for len(myMap) > 0 {
		var maxHeight int = 0
		var nume string = ""
		for name, height := range myMap {
			if height > maxHeight {
				maxHeight = height
				nume = name
			}
		}
		fmt.Println(nume)
		delete(myMap, nume)
	}
}
