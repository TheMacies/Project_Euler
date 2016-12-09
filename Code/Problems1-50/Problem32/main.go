package main

import "fmt"

func powerSet(elems []int) [][]int {
	if len(elems) == 1 {
		var set [][]int
		var emptySet []int
		elemSet := []int{elems[0]}
		set = append(set, emptySet, elemSet)
		return set
	}
	currentElem := elems[0]
	elems = elems[1:]

	subPowerSet := powerSet(elems)
	setWithCurrent := make([][]int, len(subPowerSet))
	setWithoutCurrent := make([][]int, len(subPowerSet))
	copy(setWithCurrent, subPowerSet)
	copy(setWithoutCurrent, subPowerSet)

	for i := range setWithCurrent {
		setWithCurrent[i] = append(setWithCurrent[i], currentElem)
	}
	return append(setWithCurrent, setWithoutCurrent...)
}

func combinations(elems []int) [][]int {
	currentCombs := make([][]int, 1)
	currentCombs[0] = make([]int, len(elems))

	for i := range elems {
		tempCombs := currentCombs
		currentCombs = make([][]int, 0)
		for j := range tempCombs {
			for k := 0; k < len(elems); k++ {
				tempComb := make([]int, len(elems))
				copy(tempComb, tempCombs[j])
				if tempComb[k] == 0 {
					tempComb[k] = elems[i]
					currentCombs = append(currentCombs, tempComb)
				}
			}
		}
	}
	return currentCombs
}

func arrayToNumber(comb []int) int {
	sum := 0
	for i := range comb {
		sum *= 10
		sum += comb[len(comb)-1-i]
	}
	return sum
}

func numberToArray(num int) []int {
	var comb []int
	for num > 0 {
		comb = append(comb, num%10)
		num = num / 10
	}
	return comb
}

func arrayIsCorrect(arr []int) bool {
	count := make([]int, 10)
	for i := range arr {
		count[arr[i]]++
		if count[arr[i]] > 1 {
			return false
		}
	}
	if count[0] != 0 {
		return false
	}
	return true
}

func findLeftovers(arr []int) []int {
	var leftovers []int
	for i := 1; i <= 9; i++ {
		exists := false
		for j := range arr {
			if i == arr[j] {
				exists = true
				break
			}
		}
		if !exists {
			leftovers = append(leftovers, i)
		}
	}
	return leftovers
}

func checkIfMakes9(a, b, c []int) bool {
	count := make([]int, 10)
	for _, val := range a {
		count[val]++
	}
	for _, val := range b {
		count[val]++
	}
	for _, val := range c {
		count[val]++
	}
	for i := 1; i < 10; i++ {
		if count[i] != 1 {
			return false
		}
	}
	return true
}

func sumUnique(triplets [][]int) int {
	sum := 0
	for i := range triplets {
		unique := true
		for k := 0; k < i; k++ {
			if triplets[i][2] == triplets[k][2] {
				unique = false
				break
			}
		}
		if unique {
			sum += triplets[i][2]
		}
	}
	return sum
}

func main() {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	pwrSet := powerSet(digits)
	var result [][]int
	// Remove all numbers of 5 digits of more since a product would have 5 digits. 5+5 > 9
	for {
		clean := true
		for i := range pwrSet {
			if len(pwrSet[i]) >= 5 || len(pwrSet[i]) == 0 {
				pwrSet = append(pwrSet[:i], pwrSet[i+1:]...)
				clean = false
				break
			}
		}
		if clean {
			break
		}
	}

	for i := range pwrSet {
		remainingDigits := findLeftovers(pwrSet[i])
		remainPwrSet := powerSet(remainingDigits)
		leftCombinations := combinations(pwrSet[i])
		for u := range remainPwrSet {
			rightCombinations := combinations(remainPwrSet[u])
			for j := range leftCombinations {
				for k := range rightCombinations {
					mul := numberToArray(arrayToNumber(leftCombinations[j]) * arrayToNumber(rightCombinations[k]))
					if arrayIsCorrect(mul) {
						if checkIfMakes9(leftCombinations[j], rightCombinations[k], mul) {
							var triplet []int
							triplet = append(triplet,
								arrayToNumber(leftCombinations[j]),
								arrayToNumber(rightCombinations[k]),
								arrayToNumber(mul))
							result = append(result, triplet)
						}
					}
				}
			}
		}

	}
	fmt.Println(sumUnique(result))
}
