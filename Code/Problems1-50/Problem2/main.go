package main

import (
	"fmt"
)

func fibSum(a, b, max int) int {
	if a+b > max {
		return 0
	}
	if (a+b)%2 == 1 {
		return fibSum(b, a+b, max)
	}
	return (a + b + fibSum(b, a+b, max))
}

func main() {
	a := fibSum(0, 1, 4000000)
	fmt.Print(a)
}
