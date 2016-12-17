package main

import (
	"fmt"
	"math"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func sumExists(number int) bool {
	for i := 2; i < number; i++ {
		if primesLookup[i] {
			num := number - i
			num = num / 2
			sq := int(math.Sqrt(float64(num)))
			if sq*sq == num {
				return true
			}
		}
	}
	return false
}

var primesLookup []bool

func main() {
	odd := 3
	primesLookup = kit.GeneratePrimesLookup(10000)
	for ; ; odd += 2 {
		if !primesLookup[odd] && !sumExists(odd) {
			break
		}
	}
	fmt.Println(odd)
}
