package main

import (
	"fmt"
	"time"
)

func countHexagons(l int) int {
	sum := 0
	for j := 3; j <= l; j += 3 {
		sum += j * (1 + l - j + 1) * (l - j + 1) / 6
	}
	return sum
}

func main() {
	sum := 0
	for i := 3; i <= 12345; i++ {
		sum += countHexagons(i)
	}
	fmt.Println(sum)
}
