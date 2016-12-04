package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).SetInt64(1)
	b := new(big.Int).SetInt64(1)

	i := 2
	for len(b.String()) < 1000 {
		c := b
		b = new(big.Int).Add(a, b)
		a = c
		i++
	}
	fmt.Println(i)
}
