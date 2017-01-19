package main

import (
	"fmt"
	"math/big"
	"regexp"
)

func main() {
	re := regexp.MustCompile("1\\d2\\d3\\d4\\d5\\d6\\d7\\d8\\d9\\d0")
	two := big.NewInt(2)
	i := int64(1010101030)

	//Only possible numbers are those which's end is 70 or 30 because last digit is 0 and third last is 9

	for endEquals70 := false; i < 1389026624; endEquals70 = !endEquals70 {

		num := big.NewInt(i)
		num.Exp(num, two, nil)

		if re.Match([]byte(num.String())) {
			fmt.Println(i)
			return
		}
		if endEquals70 {
			i += 60
		} else {
			i += 40
		}
	}
}
