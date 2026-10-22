package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	intSlice := []float32{3.14, 7.5, 2.0, 9.81, 4.25, 6.0, 1.618, 8.33, 5.5, 0.75}
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		numar, _ := strconv.Atoi(input)
		intSlice = append(intSlice[0:(numar)], intSlice[numar+1:len(intSlice)]...)
	}
	fmt.Println(intSlice)
}
