package main

import "fmt"

func smartPower(number int) int {
	res := number
	for i := 1; i < number; i++ {
		res = res * number
		res %= 10000000000
	}
	return res
}

func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += smartPower(i)
		sum %= 10000000000
	}
	fmt.Println(sum)
}
