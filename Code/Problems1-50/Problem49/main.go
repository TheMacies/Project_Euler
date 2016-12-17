package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

var primesLookup []bool

func countDigits(num int) []int {
	counter := make([]int, 10)
	for num > 0 {
		counter[num%10]++
		num /= 10
	}
	return counter
}

func main() {
	primesLookup = kit.GeneratePrimesLookup(20000)
	var primesList []int
	for i := 1488; i <= 9999; i++ {
		if primesLookup[i] {
			primesList = append(primesList, i)
		}
	}
	for i := 0; i < len(primesList); i++ {
		for j := i + 1; j < len(primesList); j++ {
			dif := primesList[j] - primesList[i]
			if !primesLookup[primesList[j]+dif] {
				continue
			}
			digs1 := countDigits(primesList[i])
			digs2 := countDigits(primesList[j])
			digs3 := countDigits(primesList[j] + dif)
			isAnswer := true
			for k := 0; k < 10; k++ {
				if digs1[k] != digs2[k] {
					isAnswer = false
					break
				}
				if digs2[k] != digs3[k] {
					isAnswer = false
					break
				}
			}
			if isAnswer {
				fmt.Println(kit.ConcatNumbers(primesList[i], primesList[j], primesList[j]+dif))
				return
			}
		}
	}
}
