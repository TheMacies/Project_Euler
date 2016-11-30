package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if number == 1 {
		return false
	}
	for divisor := 2; divisor < number; divisor++ {
		if (number % divisor) == 0 {
			return false
		}
	}
	return true
}

func main() {
	number := 600851475143

	for i := 2; i <= number; i++ {
		if isPrime(i) && (number%i) == 0 {
			for (number % i) == 0 {
				number = number / i
			}
			fmt.Println(i)
		}
	}
}
