package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

const (
	biggestNumber = 987654321
)

func main() {

	largest := 0
	number := 1
	for {
		n := 1
		for i := 1; ; i++ {
			if kit.CountDigits(number)*(n+1) > 9 {
				break
			}
			n++
		}
		if n == 1 {
			break
		}

		concatenated := 0
		for k := 1; k <= n; k++ {
			concatenated = kit.ConcatNumbers(concatenated, k*number)
		}
		if kit.NPandigital(concatenated, 9) {
			if concatenated > largest {
				largest = concatenated
			}
		}
		number++
	}
	fmt.Println(largest)
}
