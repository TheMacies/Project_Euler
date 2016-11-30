package main

import "fmt"

func main() {
	triplet := 1000
	for a := 0; a < triplet; a++ {
		for b := 0; b < a; b++ {
			for c := 0; c < b; c++ {
				if (c*c+b*b == a*a) && (a+b+c == triplet) {
					fmt.Println(a)
					fmt.Println(b)
					fmt.Println(c)
					fmt.Println(a * b * c)
				}
			}
		}
	}

}
