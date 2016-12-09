package main

import (
	"fmt"
	"sort"
)

func getNext(number []string) []string {
	length := len(number)
	highIndex := length + 1
	closeIndex := 0

	for i := 1; i <= length; i++ {
		if closeIndex != i && i > 1 && number[length-i] < number[length-i+1] {
			continue
		}
		for j := i + 1; j <= length; j++ {
			if number[length-i] > number[length-j] {
				if highIndex > j {
					highIndex = j
					closeIndex = i
				}
			}
		}
	}
	if highIndex == length+1 {
		return nil
	}
	tmp := number[length-closeIndex]
	for k := closeIndex; k < highIndex; k++ {
		number[length-k] = number[length-k-1]
	}
	number[length-highIndex] = tmp
	sort.Strings(number[length-highIndex+1:])
	return number
}

func main() {
	var numbers []string
	numbers = append(numbers, "0", "1", "2", "3", "4", "5", "6", "7", "8", "9")
	permutationNumber := 1
	for numbers != nil {
		numbers = getNext(numbers)
		permutationNumber++
		if permutationNumber == 1000000 {
			break
		}
	}
	fmt.Println(numbers)
}
