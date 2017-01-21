package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func main() {
	number := 600851475143
	primes := kit.GeneratePrimesLookup(7000)
	for i := 2; ; i++ {
		if primes[i] && (number%i) == 0 {
			for (number % i) == 0 {
				number = number / i
			}
			if number == 1 {
				fmt.Println(i)
				break
			}
		}
	}
}
