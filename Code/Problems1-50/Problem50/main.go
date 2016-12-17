package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

var primesLookup []bool

func main() {
	primesLookup = kit.GeneratePrimesLookup(1000000)
	max := 0
	longest := 0
	for j := 2; j < 1000000; j++ {
		sum := 0
		added := 0
		for i := j; i < 1000000 && sum < 1000000; i++ {
			if primesLookup[sum] && added > longest && added > 1 {
				max = sum
				longest = added
			}
			if primesLookup[i] {
				sum += i
				added++
			}
		}
	}
	fmt.Println(max)
}
