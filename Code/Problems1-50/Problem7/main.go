package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func main() {
	lookup := kit.GeneratePrimesLookup(1000000)
	primes := 0
	for i := 0; i < 1000000; i++ {
		if lookup[i] {
			primes++
			if primes == 10001 {
				fmt.Println(i)
				return
			}
		}
	}
}
