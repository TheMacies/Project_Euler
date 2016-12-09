package main

import "fmt"

func main() {
	factorials := make([]int, 10)

	factorials[0] = 1
	factorials[1] = 1

	for i := 2; i < 10; i++ {
		factorials[i] = factorials[i-1] * i
	}

	sum := 0

	for i := 3; i < 3000000; i++ {
		temp := i
		factorialsSum := 0
		for temp > 0 {
			factorialsSum += factorials[temp%10]
			temp /= 10
		}
		if i == factorialsSum {
			sum += i
		}
	}
	fmt.Println(sum)
}
