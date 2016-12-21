package main

import "fmt"

var preCalc [][]int

func countCombs(sum, maxNumberToUse int) int {
	if preCalc[sum][maxNumberToUse] != -1 {
		return preCalc[sum][maxNumberToUse]
	}

	if sum == 1 || maxNumberToUse == 1 {
		preCalc[sum][maxNumberToUse] = 1

		return 1
	}

	if sum < 1 || maxNumberToUse < 1 {
		preCalc[sum][maxNumberToUse] = 1

		return 1
	}
	combs := 0

	for i := 1; i <= maxNumberToUse; i++ {
		if i > sum {
			break
		}
		combs += countCombs(sum-i, i)
	}

	preCalc[sum][maxNumberToUse] = combs
	return combs
}

func main() {
	preCalc = make([][]int, 101)
	for i := 0; i <= 100; i++ {
		preCalc[i] = make([]int, 101)
		for j := 0; j <= 100; j++ {
			preCalc[i][j] = -1
		}
	}
	combinations := countCombs(101, 101)
	fmt.Println(combinations - 1)
}
