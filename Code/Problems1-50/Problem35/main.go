package main

import "fmt"

func countDigits(num int) int {
	dig := 0
	for num > 0 {
		dig++
		num /= 10
	}

	return dig
}

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

func generateRotations(i int) []int {
	var prem []int
	n := countDigits(i)
	for k := 0; k < n; k++ {
		last := i % 10
		i /= 10

		for j := 0; j < n-1; j++ {
			last *= 10
		}
		i += last
		prem = append(prem, i)
	}
	return prem
}

func main() {
	primesLookup := eratostenesSieve(1000000)
	var circPrimes []int
	circPrimes = append(circPrimes, 2)
	for i := 3; i < 1000000; i += 2 {
		circ := true
		if !primesLookup[i] {
			continue
		}

		permutations := generateRotations(i)
		for j := range permutations {
			if !primesLookup[permutations[j]] {
				circ = false
				break
			}
		}
		if circ {
			circPrimes = append(circPrimes, i)
		}
	}
	fmt.Println(circPrimes, len(circPrimes))
}
