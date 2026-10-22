package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{7, 2, 5, 4, 7, 9, 5, 3, 5}
	const n = len(intArr)
	var swapped bool
	for i := 0; i <= n-1; i++ {
		swapped = false
		for j := 0; j <= n-i-2; j++ {
			if intArr[j] > intArr[j+1] {
				var aux int32
				aux = intArr[j]
				intArr[j] = intArr[j+1]
				intArr[j+1] = aux
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Println(intArr)
}
