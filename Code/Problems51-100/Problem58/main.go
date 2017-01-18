package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func main() {
	lookup := kit.GeneratePrimesLookup(700000000)
	primes := 0.0
	allNums := 1.0

	lastDiagNumber := 1
	for side := 3; ; side += 2 {
		lastDiagNumber += side - 1
		if lookup[lastDiagNumber] {
			primes += 1.0
		}
		lastDiagNumber += side - 1
		if lookup[lastDiagNumber] {
			primes += 1.0
		}
		lastDiagNumber += side - 1
		if lookup[lastDiagNumber] {
			primes += 1.0
		}
		lastDiagNumber += side - 1
		if lookup[lastDiagNumber] {
			primes += 1.0
		}

		allNums += 4.0
		if primes/allNums < 0.1 {
			fmt.Println(side)
			return
		}
	}

}
