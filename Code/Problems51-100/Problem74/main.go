package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

var counted int
var factorials []int

func countNonRepeatingTerms(number int) int {
	n := number
	var currentTerms []int
	terms := 1
	for {
		n = getNext(n)
		for _, val := range currentTerms {
			if val == n {
				return terms
			}
		}
		currentTerms = append(currentTerms, n)
		terms++
	}
}

func getNext(i int) int {
	sum := 0
	for _, d := range kit.NumberToArray(i) {
		sum += factorials[d]
	}
	return sum
}

func main() {
	factorials = make([]int, 10)
	factorials[0] = 1
	for i := 1; i < 10; i++ {
		factorials[i] = i * factorials[i-1]
	}

	for i := 1; i <= 999999; i++ {
		if countNonRepeatingTerms(i) == 60 {
			counted++
		}
	}

	fmt.Println(counted)
}
