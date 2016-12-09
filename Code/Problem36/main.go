package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeBase2(number int) bool {
	var bit []int
	if number == 0 {
		return true
	}
	for number > 0 {
		bit = append(bit, number%2)
		number /= 2
	}
	for i := 0; i <= len(bit)/2; i++ {
		if bit[i] != bit[len(bit)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	for i := 1; i < 1000000; i++ {
		numString := strconv.Itoa(i)
		if isPalindrome(numString) && isPalindromeBase2(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
