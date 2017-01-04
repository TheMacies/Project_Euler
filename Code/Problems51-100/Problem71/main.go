package main

import (
	"fmt"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

const (
	baseValue = 3.0 / 7.0
)

type fractionsGroup struct {
	denumerator              int
	biggestPossibleNumerator int
}

func determineEndPoint(group *fractionsGroup) {
	a := 1
	b := group.denumerator - 1
	for {
		if a == b {
			group.biggestPossibleNumerator = a
			return
		}
		numerator := (a + b) / 2
		dif := float32(numerator)/float32(group.denumerator) - baseValue
		if dif == 0 {
			group.biggestPossibleNumerator = numerator - 1
			return
		}
		if dif > 0 {
			if float32(numerator-1)/float32(group.denumerator)-baseValue < 0 {
				group.biggestPossibleNumerator = numerator - 1
				return
			}
			b = numerator - 1
		} else {
			if float32(numerator+1)/float32(group.denumerator)-baseValue > 0 {
				group.biggestPossibleNumerator = numerator
				return
			}
			a = numerator + 1
		}
	}
}

var lookup []bool

func main() {
	lookup = kit.GeneratePrimesLookup(1000001)
	groups := make([]fractionsGroup, 1000001)
	for i := range groups {
		groups[i].denumerator = i
		if i < 2 {
			continue
		}
		determineEndPoint(&groups[i])
	}
	closest := 1.0
	closePoint := 0
	for i := 3; i <= 1000000; i++ {
		realEndPoint := groups[i].biggestPossibleNumerator
		if !lookup[i] {
			for kit.HCF(realEndPoint, i) != 1 {
				realEndPoint--
			}
		}

		if a := baseValue - float64(realEndPoint)/float64(i); a < closest {
			closest = a
			closePoint = realEndPoint
		}
	}
	fmt.Println(closePoint)
}
