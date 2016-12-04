package main

import "fmt"

const (
	tableSize = 3000000
)

var primesLookup []bool

func initializePrimes(table []bool) {
	tempTable := make([]bool, len(table))
	for i := 2; i < len(table); i++ {
		if tempTable[i] {
			continue
		}
		table[i] = true
		for j := i; j < len(table); j += i {
			tempTable[j] = true
		}
	}
}

func countPrimes(a, b int) int {
	n := 0
	for {
		value := n*n + n*a + b
		// a is prime if and only if -a is prime
		if value < 0 {
			value = -value
		}
		if !primesLookup[value] {
			n--
			break
		}
		n++
	}
	return n
}

func main() {
	primesLookup = make([]bool, tableSize)
	initializePrimes(primesLookup)

	maxB := 0
	maxA := 0
	maxN := 0

	for b := 1000; b > maxN; b-- {
		for a := -1000; a <= 1000; a++ {
			ct := countPrimes(a, b)
			if ct > maxN {
				maxN = ct
				maxA = a
				maxB = b
			}
		}
	}

	for b := -maxN - 1; b >= -1000; b-- {
		for a := -1000; a <= 1000; a++ {
			ct := countPrimes(a, b)
			if ct > maxN {
				maxN = ct
				maxA = a
				maxB = b
			}
		}
	}
	fmt.Println(maxN, maxA, maxB, maxA*maxB)
}
