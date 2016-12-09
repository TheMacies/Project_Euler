package main

import (
	"fmt"
	"math"
)

func main() {
	var divisors [21]int
	for i := 20; i > 0; i-- {
		var currentDivisors [21]int
		current := i
		for divisor := 2; current > 1; divisor++ {
			for current%divisor == 0 {
				current = current / divisor
				currentDivisors[divisor]++
			}
		}
		for j := 1; j <= 20; j++ {
			divisors[j] = int(math.Max(float64(divisors[j]), float64(currentDivisors[j])))
		}
	}
	product := 1
	for i := 1; i < 21; i++ {
		for j := 0; j < divisors[i]; j++ {
			product = product * i
		}
	}
	fmt.Println(product)
}
