package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	array := [...]int{2, 7, 11, 8, 4, 9}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	numar, _ := strconv.Atoi(input)
	mmap := map[int]int{}
	for i := range array {
		compl := numar - array[i]
		val, ok := mmap[compl]
		if ok {
			fmt.Println(val, i)
		}
		mmap[array[i]] = i
	}
}
