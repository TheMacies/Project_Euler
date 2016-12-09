package main

import (
	"fmt"
	"math"
)

func getDigitCount(number int) int {
	return 1 + int(math.Floor(math.Log10(float64(number))))
}

func getNthDigit(number, n int) int {
	return (int(float64(number) / math.Pow(10, float64(n-1)))) % 10
}

func isPalindrome(number int) bool {
	digitCount := getDigitCount(number)
	for i := 1; i <= digitCount/2; i++ {
		if getNthDigit(number, i) != getNthDigit(number, digitCount-i+1) {
			return false
		}
	}
	return true
}

func main() {
	currentMax := 0
	for i := 500; i < 1000; i++ {
		for j := 500; j < 1000; j++ {
			if isPalindrome(i*j) && (i*j) > currentMax {
				currentMax = i * j
			}
		}
	}
	fmt.Println(currentMax)
	fmt.Println(getNthDigit(12345, 22))
	fmt.Println(getDigitCount(12345))
	fmt.Println(isPalindrome(990009))
}
