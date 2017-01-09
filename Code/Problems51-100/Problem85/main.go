package main

import "fmt"

var groups [][]int

const (
	s = 10000
)

func main() {
	area := 0
	nearest := 2000000
	var diff int
	groups = make([][]int, s)
	for g := range groups {
		groups[g] = make([]int, s)
	}
	for i := 1; i < s; i++ {
		groups[1][i] = ((i + 1) * i) / 2
	}
	for a := 2; a < s; a++ {
		for b := a; b < s; b++ {
			groups[a][b] = groups[a-1][b] + a*((b+1)*b)/2
			diff = 2000000 - groups[a][b]
			if diff < 0 {
				diff = -diff
			}
			if diff < nearest {
				nearest = diff
				area = a * b
			}
		}
	}
	fmt.Println(area)
}
