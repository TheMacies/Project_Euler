package main

import (
	"fmt"
	"math"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

var lookup []bool

func isCorrectGroup(currentGroup []int, primeToAdd int) bool {
	for _, p := range currentGroup {
		if !lookup[kit.ConcatNumbers(primeToAdd, p)] || !lookup[kit.ConcatNumbers(p, primeToAdd)] {
			return false
		}
	}
	return true
}

func main() {
	lookup = kit.GeneratePrimesLookup(85008501)
	var currentGroups [][]int
	var usablePrimes []int

	for i := 2; i < 8500; i++ {
		if lookup[i] {
			usablePrimes = append(usablePrimes, i)
		}
	}

	for _, i := range usablePrimes {
		newGroup := []int{i}
		currentGroups = append(currentGroups, newGroup)
	}

	for groupsSize := 1; groupsSize < 5; groupsSize++ {
		var newGroups [][]int

		for _, g := range currentGroups {
			for _, p := range usablePrimes {
				if isCorrectGroup(g, p) {
					newGroup := make([]int, len(g)+1)
					copy(newGroup, g)
					newGroup[len(g)] = p
					newGroups = append(newGroups, newGroup)
				}
			}
		}
		currentGroups = newGroups
	}
	minSum := math.MaxInt32
	for _, g := range currentGroups {
		sum := 0
		for _, val := range g {
			sum += val
		}
		if sum < minSum {
			minSum = sum
		}
	}
	fmt.Println(minSum)
}
