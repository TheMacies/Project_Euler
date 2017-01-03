package main

import "fmt"

var checked []int

func count(ch []int) int {
	sum := 0
	for i := range checked {
		if checked[i] == 1 {
			sum++
		}
	}
	return sum
}

func getNext(i int) int {
	sum := 0
	for i != 0 {
		a := i % 10
		sum += a * a
		i /= 10
	}
	return sum
}

func fill(indexes []int, val int) {
	for _, ind := range indexes {
		checked[ind] = val
	}
}

func main() {
	checked = make([]int, 20000001)
	var checkedNumbers []int
	for i := 1; i < 10000000; i++ {
		if checked[i] != 0 {
			continue
		}
		for j := i; ; j = getNext(j) {
			if j == 1 {
				checked[j] = -1
			}
			if j == 89 {
				checked[j] = 1
			}
			if checked[j] == 1 {
				fill(checkedNumbers, 1)
				checkedNumbers = make([]int, 0)
				break
			}
			if checked[j] == -1 {
				fill(checkedNumbers, -1)
				checkedNumbers = make([]int, 0)
				break
			}
			checkedNumbers = append(checkedNumbers, j)
		}
	}
	fmt.Println(count(checked))
}
