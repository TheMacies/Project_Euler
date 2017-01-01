package main

import "fmt"

func countDigitOccurance(number int) []int {
	var oc []int
	oc = make([]int, 10)
	for number > 0 {
		oc[number%10]++
		number /= 10
	}
	return oc
}

func compareTables(t1, t2 []int) bool {
	for i := range t1 {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}

func main() {
	for i := 100; ; i++ {
		occ := countDigitOccurance(i)
		if !compareTables(occ, countDigitOccurance(i*2)) {
			continue
		}
		if !compareTables(occ, countDigitOccurance(i*3)) {
			continue
		}
		if !compareTables(occ, countDigitOccurance(i*4)) {
			continue
		}
		if !compareTables(occ, countDigitOccurance(i*5)) {
			continue
		}
		if !compareTables(occ, countDigitOccurance(i*6)) {
			continue
		}
		fmt.Println(i)
		return
	}
}
