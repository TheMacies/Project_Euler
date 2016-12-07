package main

import "fmt"

func getFactorsTable(number int) []int {
	factors := make([]int, 100)
	for number > 1 {
		for i := 2; i <= number; i++ {
			if number%i == 0 {
				number = number / i
				factors[i]++
				break
			}
		}
	}
	return factors
}

func numberExists(numbersFactors [][]int, searchedNumberFactors []int) bool {
	for i := range numbersFactors {
		exists := true
		for j := 0; j < len(numbersFactors[i]); j++ {
			if searchedNumberFactors[j] != numbersFactors[i][j] {
				exists = false
				break
			}
		}
		if exists {
			return true
		}
	}
	return false
}

func multiplyArray(arr []int, multiplayer int) []int {
	for i := range arr {
		arr[i] *= multiplayer
	}
	return arr
}

func main() {
	var numberFactors [][]int
	for i := range numberFactors {
		numberFactors[i] = make([]int, 100)
	}
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			arr := multiplyArray(getFactorsTable(a), b)
			if !numberExists(numberFactors, arr) {
				numberFactors = append(numberFactors, arr)
			}
		}
	}
	fmt.Println(len(numberFactors))
}
