package main

import (
	"fmt"
	"math"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func sumDigits(number int) int {
	a := kit.NumberToArray(number)
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum
}

//Not worth to use sieve since 10^14 numbers
func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	rang := int(math.Sqrt(float64(number)))
	for i := 2; i <= rang; i++ {
		if (number % i) == 0 {
			return false
		}
	}
	return true
}

func isHarashad(number int) bool {

	return number%sumDigits(number) == 0
}

func generateTruncatableHarashads(maxDigits int) []int {
	if maxDigits == 0 {
		return nil
	}
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	har := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if maxDigits == 1 {
		return har
	}
	lastSizeNumbers := 9
	for dig := 1; dig < maxDigits; dig++ {
		var newNumbers []int
		for i := 0; i < lastSizeNumbers; i++ {
			actualNumber := har[len(har)-lastSizeNumbers+i]
			for _, d := range digits {
				if isHarashad(actualNumber*10 + d) {
					newNumbers = append(newNumbers, actualNumber*10+d)
				}
			}
		}
		lastSizeNumbers = len(newNumbers)
		har = append(har, newNumbers...)
	}
	return har
}

func createStrongPrimes(har []int) []int {
	var strongHar []int
	primeDigits := []int{1, 3, 7, 9}
	for _, h := range har {
		if !isPrime(h / sumDigits(h)) {
			continue
		}
		for _, d := range primeDigits {
			if isPrime(10*h + d) {
				strongHar = append(strongHar, 10*h+d)
			}
		}
	}
	return strongHar
}

func sumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	harshads := generateTruncatableHarashads(13)
	harshads = createStrongPrimes(harshads)
	fmt.Println(sumArray(harshads))
}
