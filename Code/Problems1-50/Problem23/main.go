package main

import "fmt"

func sumDivisors(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func isAbundant(a int) bool {
	return sumDivisors(a) > a
}

func existsPair(sum int, numbers []int) bool {
	for i, a := range numbers {
		for _, b := range numbers[i:] {
			if a+b == sum {
				return true
			}
		}
	}
	return false
}

func main() {
	sum := 0
	var abundantNumbers []int
	for i := 1; i <= 28123; i++ {
		if !existsPair(i, abundantNumbers) {
			sum += i
		}
		if isAbundant(i) {
			abundantNumbers = append(abundantNumbers, i)
		}
	}
	fmt.Println(sum)
}
