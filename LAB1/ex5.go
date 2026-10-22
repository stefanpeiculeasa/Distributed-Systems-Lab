package main

import (
	"fmt"
)

func main() {
	slice := []string{"Laboratorul", "de", "sisteme", "distribuite"}
	size := len(slice)
	loopSize := size / 2
	for i := 0; i <= loopSize-1; i++ {
		aux := slice[i]
		slice[i] = slice[size-i-1]
		slice[size-i-1] = aux
	}
	fmt.Println(slice)
}
