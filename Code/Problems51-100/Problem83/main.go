package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	kit "github.com/TheMacies/Project_Euler_Toolkit"
)

func main() {
	var matrix [80][80]int
	file, _ := ioutil.ReadFile("p083_matrix.txt")
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

			if j != 0 {
				g.AddEdge(80*i+j, 80*i+j-1, matrix[i][j-1])
			}
			if i != 79 {
				g.AddEdge(80*i+j, 80*(i+1)+j, matrix[i+1][j])
			}
			if i != 0 {
				g.AddEdge(80*i+j, 80*(i-1)+j, matrix[i-1][j])
			}
		}
	}
	path, _ := g.Dijkstra(0)
	fmt.Println(path[80*80-1] + matrix[0][0])
}
