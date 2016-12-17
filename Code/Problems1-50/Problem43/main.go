package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func main() {
	set := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	perms := kit.GeneratePermutations(set)
	sum := 0
	for i := range perms {
		current := perms[i]
		if current[0] == 0 {
			continue
		}

		if kit.ConcatNumbers(current[1], current[2], current[3])%2 != 0 {
			continue
		}
		if kit.ConcatNumbers(current[2], current[3], current[4])%3 != 0 {
			continue
		}
		if kit.ConcatNumbers(current[3], current[4], current[5])%5 != 0 {
			continue
		}
		if kit.ConcatNumbers(current[4], current[5], current[6])%7 != 0 {
			continue
		}
		if kit.ConcatNumbers(current[5], current[6], current[7])%11 != 0 {
			continue
		}
		if kit.ConcatNumbers(current[6], current[7], current[8])%13 != 0 {
			continue
		}
		if kit.ConcatNumbers(current[7], current[8], current[9])%17 != 0 {
			continue
		}
		sum += kit.ArrayToNumber(current)
	}
	fmt.Println(sum)
}
