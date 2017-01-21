package main

import "fmt"

func main() {
	var sum, sumOfSquares int64

	for i := int64(1); i <= 100; i++ {
		sum += i
		sumOfSquares += i * i
	}
	sum = sum * sum
	fmt.Println(sumOfSquares - sum)

}
