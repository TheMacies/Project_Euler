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
	file, _ := ioutil.ReadFile("p081_matrix.txt")
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
		}
	}
	val, err := g.Dijkstra(0, 80*80-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matrix[0][0] + val)
}
