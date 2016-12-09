package main

import "fmt"

func main() {
	max := 1
	numMax := 1

	for i := 2; i < 1000000; i++ {
		current := 1
		number := i
		for number != 1 {
			if number%2 == 0 {
				number = number / 2
			} else {
				number = 3*number + 1
			}
			current++
		}
		if current > max {
			max = current
			numMax = i
		}
	}
	fmt.Println(numMax)
}
