package main

import (
	"fmt"
	"math/big"
)

func power(a, b int) *big.Int {
	res := big.NewInt(1)
	for i := 0; i < b; i++ {
		res.Mul(res, big.NewInt(int64(a)))
	}
	return res
}

func main() {
	counter := 0
	for i := 1; ; i++ {
		for j := 1; ; j++ {
			l := len(power(j, i).String())
			if l < i {
				if j == 9 {
					fmt.Println(counter)
					return
				}
				continue
			}
			if l > i {
				break
			}
			counter++
		}
	}

}
