package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func main() {
	max := 1
	lookup := kit.GeneratePrimesLookup(1000)
	for i, val := range lookup {
		if val {
			if i*max < 1000000 {
				max *= i
			} else {
				break
			}
		}
	}
	fmt.Println(max)
}
