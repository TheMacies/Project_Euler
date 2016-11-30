package main

import (
	"fmt"
)

func main() {
	const tabSize int = 2000000
	var primes [tabSize]int
	primes[0] = 1
	primes[1] = 1
	for i := 2; i < tabSize; i++ {
		if primes[i] == 0 {
			for j := i * 2; j < tabSize; j += i {
				primes[j] = 1
			}
		}
	}
	sum := 0
	for i := 0; i < tabSize; i++ {
		if primes[i] == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
