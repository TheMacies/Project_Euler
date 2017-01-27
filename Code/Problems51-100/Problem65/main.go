package main

import (
	"fmt"
	"math/big"
)

type fraction struct {
	numerator, denom *big.Int
}

func sum(val string) int {
	sum := 0
	for _, r := range val {
		sum += int(r) - 48
	}
	return sum
}

func main() {
	convSteps := 99
	f := fraction{big.NewInt(1), big.NewInt(1)}

	if convSteps%3 == 2 {
		mul := (convSteps + 1) / 3
		mul *= 2
		f.numerator = big.NewInt(int64(mul))
	}

	for i := convSteps - 1; i >= 1; i-- {
		f.denom, f.numerator = f.numerator, f.denom
		mul := 1
		if i%3 == 2 {
			mul = (i + 1) / 3
			mul *= 2
		}
		f.numerator.Add(f.numerator, big.NewInt(0).Mul(big.NewInt(int64(mul)), f.denom))
	}

	f.denom, f.numerator = f.numerator, f.denom
	f.numerator.Add(f.numerator, big.NewInt(0).Mul(big.NewInt(int64(2)), f.denom))
	fmt.Println(sum(f.numerator.String()))

}
