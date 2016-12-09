package main

import "fmt"

func eratostenesSieve(n int) []bool {
	primes := make([]bool, n)
	temp := make([]bool, n)
	for i := 2; i < n; i++ {
		if temp[i] == false {
			primes[i] = true
			for j := i; j < n; j += i {
				temp[j] = true
			}
		}
	}
	return primes
}

func leftTruncatable(number int, primesLookup []bool) bool {
	for number > 9 {
		number /= 10
		if !primesLookup[number] {
			return false
		}
	}
	return true
}

func rightTruncatable(number int, primesLookup []bool) bool {
	modulo := 1
	copy := number
	for copy > 9 {
		modulo *= 10
		copy /= 10
	}

	for number > 9 {
		number = number % modulo
		modulo /= 10
		if !primesLookup[number] {
			return false
		}
	}
	return true
}

func main() {
	primesLookup := eratostenesSieve(1000000)
	sum := 0
	primesFound := 0
	for i := 11; i < 1000000; i++ {
		if primesFound == 11 {
			break
		}
		if !primesLookup[i] {
			continue
		}
		if leftTruncatable(i, primesLookup) && rightTruncatable(i, primesLookup) {
			sum += i
			primesFound++
		}
	}
	fmt.Println(primesFound, sum)
}
