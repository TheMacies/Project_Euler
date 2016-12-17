package main

import (
	"fmt"
	"math/big"
)

var primesLoookup []bool

func recurringCycle(num int) int {
	pow := big.NewInt(10)
	for i := 1; i < num; i++ {
		mod := new(big.Int)
		mod.Mod(pow, big.NewInt(int64(num)))
		if mod.Cmp(big.NewInt(1)) == 0 {
			return i + 1
		}
		pow.Mul(pow, big.NewInt(10))
	}
	return 0
}

func main() {
	max := 0
	for i := 2; i < 1000; i++ {
		cycle := recurringCycle(i)
		if cycle > max {
			max = cycle
		}
	}
	fmt.Println(max)
}
