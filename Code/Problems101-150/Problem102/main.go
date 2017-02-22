package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Point struct {
	x, y float64
}

func rectangleCouldHaveOrigin(a, b, c Point) bool {
	minX, maxX, minY, maxY := a.x, a.x, a.y, a.y
	minX = min(minX, b.x)
	maxX = max(maxX, b.x)
	minY = min(minY, b.y)
	maxY = max(maxY, b.y)
	minX = min(minX, c.x)
	maxX = max(maxX, c.x)
	minY = min(minY, c.y)
	maxY = max(maxY, c.y)
	if minX > 0 || maxX < 0 || minY > 0 || maxY < 0 {
		return false
	}
	return true
}

func min(a, b float64) {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) {
	if a > b {
		return a
	}
	return b
}

func pointsMakeLineParallelToAnyAxis(a, b Point) bool {
	if a.x == b.x || a.y == b.y {
		return true
	}
	return false
}

func countParallelSides(a, b, c Point) int {
	parallel := 0
	if pointsMakeLineParallelToAnyAxis(a, b) {
		parallel++
	}

	if pointsMakeLineParallelToAnyAxis(a, c) {
		parallel++
	}

	if pointsMakeLineParallelToAnyAxis(c, b) {
		parallel++
	}
	return parallel
}

func sortPointsToMakeANearRightAngle(a, b, c Point) (Point, Point, Point) {
	if a.x == b.x && a.y == c.y {
		return a, b, c
	}

	if a.x == c.x && a.y == b.y {
		return a, b, c
	}

	if b.x == c.x && b.y == a.y {
		return b, a, c
	}

	if b.x == a.x && b.y == c.y {
		return b, a, c
	}
	return c, a, b
}

func sortPointsToMakeABParallel(a, b, c Point) (Point, Point, Point) {
	if a.x == b.x && a.y == b.y {
		return a, b, c
	}

	if a.x == c.x && a.y == c.y {
		return a, c, b
	}

	return b, c, a
}

func URRightTriangleHasOrigin(a, b, c Point) bool {
	a, b, c = sortPointsToMakeANearRightAngle(a, b, c)

	if b.x > c.x {
		b, c = c, b
	}

	if b.x > 0 || c.x < 0 || b.y > 0 || c.y < 0 {
		return false
	}

	deltaY := (c.y - b.y) / (c.x - b.x)

	if b.y+deltaY*(-b.x) < 0 {
		return true
	}
	return false
}

func ULRightTriangleHasOrigin(a, b, c Point) bool {
	a, b, c = sortPointsToMakeANearRightAngle(a, b, c)

	if b.x > c.x {
		b, c = c, b
	}

	if b.x > 0 || c.x < 0 || b.y < 0 || c.y > 0 {
		return false
	}

	deltaY := (c.y - b.y) / (c.x - b.x)

	if b.y+deltaY*(-b.x) < 0 {
		return true
	}
	return false
}

func DRRightTriangleHasOrigin(a, b, c Point) bool {
	a, b, c = sortPointsToMakeANearRightAngle(a, b, c)

	if b.x > c.x {
		b, c = c, b
	}

	if b.x > 0 || c.x < 0 || b.y > 0 || c.y < 0 {
		return false
	}

	deltaY := (c.y - b.y) / (c.x - b.x)

	if b.y+deltaY*(-b.x) > 0 {
		return true
	}
	return false
}

func DLRightTriangleHasOrigin(a, b, c Point) bool {
	a, b, c = sortPointsToMakeANearRightAngle(a, b, c)

	if b.x > c.x {
		b, c = c, b
	}

	if b.x > 0 || c.x < 0 || b.y < 0 || c.y > 0 {
		return false
	}

	deltaY := (c.y - b.y) / (c.x - b.x)

	if b.y+deltaY*(-b.x) > 0 {
		return true
	}
	return false
}

func triangleHasOrigin(a, b, c Point) bool {
	if !rectangleCouldHaveOrigin(a, b, c) {
		return false
	}
	parallel := countParallelSides(a, b, c)
	if parallel == 2 {
		return rightTriangleHasOrigin(a, b, c)
	}

	if parallel == 1 {
		return parallelTriangleHasOrigin(a, b, c)
	}
	return regularTriangleHasOrigin(a, b, c)
}

func rightTriangleHasOrigin(a, b, c Point) bool {
	if !rectangleCouldHaveOrigin(a, b, c) {
		return false
	}
	a, b, c = sortPointsToMakeANearRightAngle(a, b, c)
	if a.x < b.x || a.x < c.x {
		if a.y < b.y || a.y < c.y {
			return DLRightTriangleHasOrigin(a, b, c)
		}
		return ULRightTriangleHasOrigin(a, b, c)
	}

	if a.y < b.y || a.y < c.y {
		return DRRightTriangleHasOrigin(a, b, c)
	}
	return URRightTriangleHasOrigin(a, b, c)
}

func parallelTriangleHasOrigin(a, b, c Point) bool {
	a, b, c = sortPointsToMakeABParallel(a, b, c)
	if a.x == b.x {
		if b.y > a.y {
			a, b = b, a
		}
		if c.x > a.x {
			if c.y > a.y {
				return !ULRightTriangleHasOrigin(a, Point{a.x, c.y}, c) && !DRRightTriangleHasOrigin(Point{c.x, b.y}, b, c)
			} else if c.y < b.y {
				return !DLRightTriangleHasOrigin(b, Point{b.x, c.y}, c) && !URRightTriangleHasOrigin(Point{c.x, a.y}, a, c)
			} else {
				return ULRightTriangleHasOrigin(a, Point{a.x, c.y}, c) || DRRightTriangleHasOrigin(Point{b.x, c.y}, b, c)
			}
		} else {
			if c.y > a.y {
				return !URRightTriangleHasOrigin(a, Point{a.x, c.y}, c) && !DLRightTriangleHasOrigin(Point{c.x, b.y}, b, c)
			} else if c.y < b.y {
				return !DRRightTriangleHasOrigin(b, Point{b.x, c.y}, c) && !ULRightTriangleHasOrigin(Point{c.x, a.y}, a, c)
			} else {
				return ULRightTriangleHasOrigin(a, Point{a.x, c.y}, c) || DRRightTriangleHasOrigin(Point{b.x, c.y}, b, c)
			}
		}
	} else {

		if a.x > b.x {
			a, b = b, a
		}
		if c.y > a.y {
			if c.x > a.x {
				return !ULRightTriangleHasOrigin(a, Point{a.x, c.y}, c) && !DRRightTriangleHasOrigin(Point{c.x, b.y}, b, c)
			} else if c.x < b.x {
				return !DLRightTriangleHasOrigin(b, Point{b.x, c.y}, c) && !URRightTriangleHasOrigin(Point{c.x, a.y}, a, c)
			} else {
				return ULRightTriangleHasOrigin(a, Point{a.x, c.y}, c) || DRRightTriangleHasOrigin(Point{b.x, c.y}, b, c)
			}
		}
	}
}

func main() {
	fileText, err := ioutil.ReadFile("p102_triangles.txt")
	if err != nil {
		log.Fatal(err)
	}
	count := 0

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

		if triangleHasOrigin(a, b, c) {
			count++
		}
	}
	fmt.Println(count)
}
