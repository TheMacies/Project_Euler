package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).SetInt64(1)
	for i := 0; i < 1000; i++ {
		a = a.Mul(a, new(big.Int).SetInt64(2))
	}
	res := a.String()
	sum := 0
	for i := 0; i < len(res); i++ {
		sum += int(res[i]) - 48
	}
	fmt.Println(sum)
}
