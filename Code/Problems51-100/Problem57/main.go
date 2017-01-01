package main

import (
	"fmt"
	"math/big"
)

type fract struct {
	num, denom *big.Int
}

func main() {
	counter := 0
	fracts := make([]fract, 1000)
	fracts[0].num = big.NewInt(1)
	fracts[0].denom = big.NewInt(2)
	for i := 1; i < 1000; i++ {
		fracts[i].num = big.NewInt(0)
		fracts[i].num.Add(fracts[i-1].num, fracts[i-1].denom)
		fracts[i].num.Add(fracts[i].num, fracts[i-1].denom)
		fracts[i].denom = big.NewInt(0).Add(big.NewInt(0), fracts[i-1].denom)
		fracts[i].denom, fracts[i].num = fracts[i].num, fracts[i].denom
	}
	for i := 0; i < 1000; i++ {
		fracts[i].num.Add(fracts[i].num, fracts[i].denom)
		if len(fracts[i].num.String()) > len(fracts[i].denom.String()) {
			counter++
		}
	}
	fmt.Println(counter)
}
