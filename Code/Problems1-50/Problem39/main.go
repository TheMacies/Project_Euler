package main

import (
	"fmt"
	"math"
)

func countSolutions(p int) int {
	solutions := 0

	if p%2 == 1 {
		return 0
	}

	for i := 1; i < p; i++ {
		for j := i + 1; j < p; j++ {
			c := int(math.Sqrt(float64(i*i + j*j)))
			if c*c != i*i+j*j {
				continue
			}
			if i+j+c == p {
				solutions++
			}
		}
	}
	return solutions
}

func main() {
	maxSolution := 0
	maxP := 0
	for p := 1; p <= 1000; p++ {
		sol := countSolutions(p)
		if sol > maxSolution {
			maxSolution = sol
			maxP = p
		}
	}
	fmt.Println(maxSolution, maxP)
}
