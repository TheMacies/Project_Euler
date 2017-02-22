package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y float64
}

func area(a, b, c Point) float64 {
	return math.Abs(a.x*(b.y-c.y)+b.x*(c.y-a.y)+c.x*(a.y-b.y)) / 2.0
}

func main() {
	fileText, err := ioutil.ReadFile("p102_triangles.txt")
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	p := Point{0, 0}

	lines := strings.Split(string(fileText), "\n")
	for _, line := range lines[:len(lines)-1] {

		values := strings.Split(line, ",")
		vals := [6]int{}
		for i := range vals {
			vals[i], err = strconv.Atoi(values[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		a, b, c :=
			Point{float64(vals[0]), float64(vals[1])},
			Point{float64(vals[2]), float64(vals[3])},
			Point{float64(vals[4]), float64(vals[5])}

		if area(a, b, c) == area(a, b, p)+area(a, c, p)+area(b, c, p) {
			count++
		}
	}
	fmt.Println(count)
}
