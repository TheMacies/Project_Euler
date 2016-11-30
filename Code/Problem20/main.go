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
	a := new(big.Int).SetInt64(100)
	fact := factorial(a).String()
	sum := 0
	for i := 0; i < len(fact); i++ {
		sum += int(fact[i]) - 48 //ASCII to numver conv
	}
	fmt.Println(sum)

}
