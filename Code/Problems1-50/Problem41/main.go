package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func main() {
	primeLookup := kit.GeneratePrimesLookup(987654322)
	max := 0
	var set []int
	set = append(set, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 9; i >= 1; i-- {
		perms := kit.GeneratePermutations(set[:i])
		for j := range perms {
			number := kit.ArrayToNumber(perms[j])
			if primeLookup[number] {
				if number > max {
					max = number
				}
			}
		}
		if max > 0 {
			fmt.Println(max)
			return
		}
	}
}
