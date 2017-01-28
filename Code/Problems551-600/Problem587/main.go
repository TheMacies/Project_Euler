package main

import (
	"fmt"
	"math"
)

var lArea float64

//Numerical integration of the area using rectangle quadrature
func calculateArea(n float64) float64 {
	a := 1.0 / n
	//Point in which function changes from linear to  circle-like
	changePoint := (2*a + 2.0 - math.Sqrt(8.0*a)) / (2.0*a*a + 2.0)
	f := func(x float64) float64 {
		if x >= changePoint {
			return 1.0 - math.Sqrt(2.0*x-x*x)
		} else {
			return a * x
		}
	}
	sum := 0.0
	h := 0.001
	for i := 0.0; i < 1.0; i += h {
		sum += h / 2 * (f(i) + f(i+h))
	}
	return sum
}

func main() {
	// Radius of circle does not matter that is why i picked r = 1
	lArea = (4.0 - math.Pi) / 4.0
	//This part is a bit complicated
	//All that it does is searching for n in some faster way. Instead of going through all n's
	//I instead look for thousands first then for houndreds (and so on) so that number of iterations is strongly decreased
	thousands := 0.0
	for i := 1.0; ; i += 1000.0 {
		if calculateArea(i)/lArea < 0.001 {

			thousands--
			break
		}
		thousands += 1.0
	}
	houndreds := 0.0
	for i := 1000.0 * thousands; ; i += 100.0 {
		if calculateArea(i)/lArea < 0.001 {

			houndreds--
			break
		}
		houndreds++
	}
	tens := 0.0
	for i := 1000.0*thousands + 100.0*houndreds; ; i += 10.0 {
		if calculateArea(i)/lArea < 0.001 {

			tens--
			break
		}
		tens++
	}
	for i := 1000.0*thousands + 100.0*houndreds + 10.0*tens; ; i += 1.0 {
		if calculateArea(i)/lArea < 0.001 {
			fmt.Println(i)
			break
		}
	}
}
