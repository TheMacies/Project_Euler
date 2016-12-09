package main

import (
	"fmt"
	"strconv"
)

func createSeq(m int) string {
	var seq string
	num := 1
	seq += "0"
	for len(seq) < m+1 {
		seq += strconv.Itoa(num)
		num++
	}
	return seq
}

func main() {
	fract := createSeq(1000000)
	d1 := int(fract[1]) - 48
	d2 := int(fract[10]) - 48
	d3 := int(fract[100]) - 48
	d4 := int(fract[1000]) - 48
	d5 := int(fract[10000]) - 48
	d6 := int(fract[100000]) - 48
	d7 := int(fract[1000000]) - 48
	fmt.Println(d1 * d2 * d3 * d4 * d5 * d6 * d7)
}
