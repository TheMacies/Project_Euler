package main

import "fmt"

func powerOfFive(n int) int {
	return n * n * n * n * n
}

func main() {
	sum := 0

	for i := 10; i < 1000000; i++ {
		a := i
		digitsPowersSum := 0
		for a > 0 {
			digitsPowersSum += powerOfFive(a % 10)
			a /= 10
		}
		if i == digitsPowersSum {
			sum += i
		}
	}

	fmt.Println(sum)
}
