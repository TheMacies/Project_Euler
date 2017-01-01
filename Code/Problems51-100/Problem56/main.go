package main

import (
	"fmt"
	"math/big"
)

func sumDigits(number string) int {
	sum := 0
	for i := range number {
		sum += int(number[i]) - 48
	}
	return sum
}

func power(a, b int) *big.Int {
	aBig := big.NewInt(int64(a))
	sum := big.NewInt(1)
	for i := 0; i < b; i++ {
		sum.Mul(sum, aBig)
	}
	return sum
}

func main() {
	max := 0

	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			val := power(a, b)
			sum := sumDigits(val.String())
			if sum > max {
				max = sum
			}
		}
	}

	fmt.Println(max)
}
