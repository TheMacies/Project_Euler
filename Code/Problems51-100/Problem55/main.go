package main

import (
	"fmt"
	"math/big"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func reverseAdd(num *big.Int) {
	a := big.NewInt(0)
	s := kit.ReverseString(num.String())
	for i := range s {
		a.Mul(a, big.NewInt(10))
		a.Add(a, big.NewInt(int64(s[i]-48)))
	}
	num.Add(a, num)
}

func main() {
	sum := 0
	for i := 1; i < 10000; i++ {
		val := big.NewInt(int64(i))
		reverseAdd(val)
		for j := 0; j < 50; j++ {
			s := val.String()
			if kit.IsPalindrome(s) {
				break
			}
			reverseAdd(val)
			if j == 49 {
				sum += 1
			}
		}
	}
	fmt.Println(sum)
}
