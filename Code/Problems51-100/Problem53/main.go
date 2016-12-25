package main

import (
	"fmt"
	"math/big"
)

func factorial(i *big.Int) *big.Int {
	if i.Cmp(big.NewInt(1)) <= 0 {
		return big.NewInt(1)
	}
	return i.Mul(i, factorial(new(big.Int).Sub(i, big.NewInt(1))))
}

func combinations(n, r int) *big.Int {
	comb := big.NewInt(1)
	comb.Mul(comb, factorial(big.NewInt(int64(n))))
	comb.Quo(comb, factorial(big.NewInt(int64(r))))
	comb.Quo(comb, factorial(big.NewInt(int64(n-r))))
	return comb
}

func main() {
	counter := 0

	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			combs := combinations(i, j)
			if combs.Cmp(big.NewInt(1000000)) > 0 {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
