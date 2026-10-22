package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		number, err := strconv.Atoi(input)
		if err == nil {
			var div int = 0
			for i := 1; i <= number; i++ {
				if number%i == 0 {
					div++
				}
			}
			if div == 2 {
				println("prim")
			} else {
				println("neprim")
			}
		}
	}
}
