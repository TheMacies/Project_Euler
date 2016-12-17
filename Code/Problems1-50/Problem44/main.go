package main

import "fmt"

func main() {
	lookup := make([]bool, 20000000)
	min := 1000000000
	var pentagons []int
	for i := 1; ; i++ {
		pen := (i * (3*i - 1)) / 2
		if pen > 10000000 {
			break
		}
		lookup[pen] = true
		pentagons = append(pentagons, pen)
	}
	for i := 0; i < len(pentagons); i++ {
		for j := i + 1; j < len(pentagons); j++ {
			sum := pentagons[i] + pentagons[j]
			diff := pentagons[j] - pentagons[i]
			if !lookup[sum] || !lookup[diff] {
				continue
			}
			if diff < min {
				min = diff
				fmt.Println(pentagons[i], pentagons[j])
			}
		}
	}
	fmt.Println(min)
}
