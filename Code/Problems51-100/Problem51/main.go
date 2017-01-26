package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

var lookup []bool
var patterns [][]int
var digits []int

func isOfEightFamily(i int) bool {
	for _, pat := range patterns {
		famSize := 0
		var temp []int
		for _, dig := range digits {
			if temp = append(temp, buildNumber(pat, i, dig)); lookup[temp[len(temp)-1]] {
				famSize++
			}
		}
		if famSize == 8 {
			for _, num := range temp {
				if num == i {
					return true
				}
			}
		}
	}
	return false
}

func buildNumber(pattern []int, num, dig int) int {
	number := 0
	l := kit.CountDigits(num)
	for i := 0; i < l; i++ {
		number *= 10
		if pattern[i] == 1 {
			number += num % 10
		} else {
			number += dig
		}
		num /= 10
	}
	return kit.ReverseNumber(number)
}

func main() {
	lookup = kit.GeneratePrimesLookup(1000000)
	digits = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	patterns = make([][]int, 10)
	patterns[0] = []int{1, 1, 1, 0, 0, 0}
	patterns[1] = []int{1, 1, 0, 1, 0, 0}
	patterns[2] = []int{1, 1, 0, 0, 1, 0}
	patterns[3] = []int{1, 1, 0, 0, 0, 1}
	patterns[4] = []int{1, 0, 1, 1, 0, 0}
	patterns[5] = []int{1, 0, 1, 0, 1, 0}
	patterns[6] = []int{1, 0, 1, 0, 0, 1}
	patterns[7] = []int{1, 0, 0, 1, 1, 0}
	patterns[8] = []int{1, 0, 0, 1, 0, 1}
	patterns[9] = []int{1, 0, 0, 0, 1, 1}
	for i := 120383; i < 1000000; i++ {
		if !lookup[i] {
			continue
		}
		if isOfEightFamily(i) {
			fmt.Println(i)
			return
		}
	}
}
