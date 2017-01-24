package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func main() {
	var matrix [80][80]int
	file, _ := ioutil.ReadFile("p082_matrix.txt")
	rows := strings.Split(string(file), "\n")
	for i, r := range rows[:len(rows)-1] { //last line is empty in the file
		for j, valString := range strings.Split(r, ",") {
			val, _ := strconv.Atoi(valString)
			matrix[i][j] = val
		}
	}
	g := kit.NewDirectedGraph()
	for i := 0; i < 80*80; i++ {
		g.AddVertex()
	}
	for i := 0; i < 80; i++ {
		for j := 0; j < 80; j++ {
			if j != 79 {
				g.AddEdge(80*i+j, 80*i+j+1, matrix[i][j+1])
			}
			if i != 79 {
				g.AddEdge(80*i+j, 80*(i+1)+j, matrix[i+1][j])
			}
			if i != 0 {
				g.AddEdge(80*i+j, 80*(i-1)+j, matrix[i-1][j])
			}
		}
	}
	min := math.MaxInt32
	for i := 0; i < 80; i++ {
		val, _ := g.Dijkstra(80 * i)

		for j := 0; j < 80; j++ {
			if matrix[i][0]+val[80*j+79] < min {
				min = val[80*j+79] + matrix[i][0]
			}
		}
	}
	fmt.Println(min)
}
