package main

import "fmt"

func main() {
	var a int64
	a = 1
	for i := 0; i < 7830457; i++ {
		a *= 2
		a %= 10000000000
	}
	a *= 28433
	a %= 10000000000
	fmt.Println(a + 1)
}
