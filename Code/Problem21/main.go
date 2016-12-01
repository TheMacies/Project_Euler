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

func main() {
	sums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		sums[i] = sumDivisors(i)
	}
	amicableSum := 0
	for i, val := range sums {
		if val < 10000 {
			if sums[val] == i && i != val {
				amicableSum += i + val
				fmt.Println(i, val)
			}
		}
	}
	fmt.Println(amicableSum / 2)
}
