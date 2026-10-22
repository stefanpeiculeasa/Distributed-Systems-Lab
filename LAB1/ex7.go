package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	slice := []string{}
	myMap := map[string]int{}
	file, _ := os.Open("ex7.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}
	for i := range slice {
		_, ok := myMap[slice[i]]
		if !ok {
			myMap[slice[i]] = 1
		} else {
			myMap[slice[i]]++
		}
	}
	fmt.Println(myMap)
}
