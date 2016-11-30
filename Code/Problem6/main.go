package main

import (
	"fmt"
	"math/big"
)

func main() {
	sumOfProducts := big.NewInt(0)
	sum := big.NewInt(0)
	for i := 1; i <= 100; i++ {
		sumOfProducts = big.NewInt(0).Add(sumOfProducts, big.NewInt(1).Mul(big.NewInt(int64(i)), big.NewInt(int64(i))))
		sum = big.NewInt(0).Add(sum, big.NewInt(int64(i)))
	}
	sumSquared := big.NewInt(1).Mul(sum, sum)
	fmt.Println(big.NewInt(0).Sub(sumSquared, sumOfProducts))
}
