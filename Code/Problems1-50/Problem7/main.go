package main

import (
	"fmt"
)

func main() {
	const tabSize int = 1000000
	var primes [tabSize]int
	primes[0] = 1
	primeCount := 0

	for i := 2; i < tabSize; i++ {
		if primes[i] == 0 {
			primeCount++
			if primeCount == 10001 {
				fmt.Println(i)
				return
			}
			for j := i * 2; j < tabSize; j += i {
				primes[j] = 1
			}
		}
	}
	fmt.Println("Not enough tabSize")
	fmt.Print("Found only ")
	fmt.Print(primeCount)
	fmt.Print(" prime numbers")
}
