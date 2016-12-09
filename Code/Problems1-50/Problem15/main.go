package main

import (
	"fmt"
	"math/big"
)

func factorial(n *big.Int) *big.Int {
	two := new(big.Int).SetInt64(2)
	if n.Cmp(two) == -1 {
		return new(big.Int).SetInt64(1)
	}
	m := new(big.Int).Set(n)
	a := factorial(n.Sub(n, new(big.Int).SetInt64(1)))
	m.Mul(m, a)
	return m
}

func main() {
	a := new(big.Int).SetInt64(40)
	b := new(big.Int).SetInt64(20)
	res1 := factorial(a)
	res2 := factorial(b)
	res1 = res1.Div(res1, res2)
	res1 = res1.Div(res1, res2)
	fmt.Println(res1)
}
