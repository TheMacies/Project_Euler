package main

import "fmt"

var grids [][]int

const (
	s = 10000
)

func main() {
	area := 0
	nearest := 2000000
	var diff int
	grids = make([][]int, s)
	for g := range grids {
		grids[g] = make([]int, s)
	}
	for i := 1; i < s; i++ {
		grids[1][i] = ((i + 1) * i) / 2
	}
	for a := 2; a < s; a++ {
		for b := a; b < s; b++ {
			grids[a][b] = grids[a-1][b] + a*((b+1)*b)/2
			diff = 2000000 - grids[a][b]
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
