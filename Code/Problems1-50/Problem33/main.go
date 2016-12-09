package main

import "fmt"

type fraction struct {
	num   int
	denom int
}

func trivial(a, b int) bool {
	if a%10 == 0 && b%10 == 0 {
		return true
	}
	if a == b {
		return true
	}
	return false
}

func curious(a, b int) bool {
	a1 := a % 10
	a2 := a / 10
	b1 := b % 10
	b2 := b / 10

	if a1 == b2 && a2*b == b1*a {
		return true
	}
	return false
}

func main() {
	var fracts []fraction
	for a := 11; a < 100; a++ {
		for b := a + 1; b < 100; b++ {
			if curious(a, b) && !trivial(a, b) {
				fracts = append(fracts, fraction{num: a, denom: b})
			}
		}
	}
	fmt.Println(fracts)
}
