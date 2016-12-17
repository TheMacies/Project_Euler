package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func numberHas4DifferentPFactors(num int) bool {
	factors := kit.GetFactors(num)
	distinctPrimes := 0
	for i := range factors {
		if factors[i] == 0 {
			continue
		}
		if !primesLookup[i] {
			return false
		}
		distinctPrimes++
		if distinctPrimes > 4 {
			return false
		}
	}
	if distinctPrimes != 4 {
		return false
	}
	return true
}

var primesLookup []bool

func main() {
	primesLookup = kit.GeneratePrimesLookup(10000000)
	for i := 2; i < 10000000; i++ {
		if !numberHas4DifferentPFactors(i) {
			continue
		}
		isAnswer := true
		for j := 1; j < 4; j++ {
			if !numberHas4DifferentPFactors(i + j) {
				isAnswer = false
				break
			}
		}
		if isAnswer {
			fmt.Println(i)
			return
		}
		i += 2
	}
}
