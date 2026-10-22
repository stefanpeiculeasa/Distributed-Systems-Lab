package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{7, 2, 5, 4, 7, 9, 5, 3, 5}
	intSlice := []int32{}
	for i := range intArr {
		if intArr[i]%2 == 0 {
			intSlice = append(intSlice, intArr[i])
		}
	}
	fmt.Println(intSlice)
}
