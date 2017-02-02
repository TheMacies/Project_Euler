package main

import "fmt"

var preCalc []int

func pentagonal(i int) int {
	return (i * (3*i - 1)) / 2
}

func findCoinNumber() int {
	preCalc := []int{1}
	n := 1
	for {
		i := 0
		penta := 1
		preCalc = append(preCalc, 0)
		for penta <= n {
			sign := 1
			if i%4 > 1 {
				sign = -1
			}
			preCalc[n] += sign * preCalc[n-penta]
			preCalc[n] %= 1000000
			i++
			j := i/2 + 1
			if i%2 == 1 {
				j *= -1
			}
			penta = pentagonal(j)
		}
		if preCalc[n] == 0 {
			return n
		}
		n++
	}
}

func main() {
	fmt.Println(findCoinNumber())
}
