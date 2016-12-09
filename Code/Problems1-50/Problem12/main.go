package main

import "fmt"

type divisor struct {
	number     int
	occurances int
}

func getPrimes(number int) []divisor {
	var primes []divisor
	for number > 1 {
		for div := 2; ; div++ {
			if number%div == 0 {
				exists := false

				for i, di := range primes {
					if di.number == div {
						primes[i].occurances++
						exists = true
						break
					}
				}
				if !exists {
					primes = append(primes, divisor{number: div, occurances: 1})
				}
				number = number / div
				break
			}
		}
	}
	return primes
}

func countPossibleDivisors(primes []divisor) int {
	combinations := 1
	for _, val := range primes {
		combinations = combinations * (val.occurances + 1)
	}
	return combinations
}

func main() {
	triangle := 1
	max := 0
	for number := 2; ; number++ {
		triangle += number
		primes := getPrimes(triangle)
		divisors := countPossibleDivisors(primes)
		if max < divisors {
			max = divisors
			fmt.Println(max)
		}
		if triangle < 0 {
			fmt.Println("overflow :()")
			break
		}
		if divisors > 500 {
			fmt.Println(triangle, " ", divisors)
			break
		}
	}

}
