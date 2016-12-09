package main

import "fmt"

func countOptions(coinsSet []int, value int) int {
	if len(coinsSet) == 1 {
		return 1
	}
	options := 0
	coinsNumber := 0
	for {
		if coinsNumber*coinsSet[0] > value {
			break
		}
		options += countOptions(coinsSet[1:], value-coinsNumber*coinsSet[0])
		coinsNumber++
	}
	return options
}

func main() {
	var coinsSet []int
	coinsSet = append(coinsSet, 200, 100, 50, 20, 10, 5, 2, 1)
	fmt.Println(countOptions(coinsSet, 200))
}
