package main

import "fmt"

func countLetters(number int) int {
	sum := 0
	if number > 99 {
		sum += 7
		sum += preCounted[number/100]
		number %= 100
		if number == 0 {
			return sum
		}
		sum += 3

	}

	if number == 0 {
		return sum
	}
	if number > 19 {
		sum += preCounted[(number/10)*10]
		if number%10 != 0 {
			sum += preCounted[number%10]
		}
	} else {
		sum += preCounted[number]
	}

	return sum
}

var preCounted map[int]int

func main() {
	preCounted = make(map[int]int)

	preCounted[1] = 3
	preCounted[2] = 3
	preCounted[3] = 5
	preCounted[4] = 4
	preCounted[5] = 4
	preCounted[6] = 3
	preCounted[7] = 5
	preCounted[8] = 5
	preCounted[9] = 4
	preCounted[10] = 3
	preCounted[11] = 6
	preCounted[12] = 6
	preCounted[13] = 8
	preCounted[14] = 8
	preCounted[15] = 7
	preCounted[16] = 7
	preCounted[17] = 9
	preCounted[18] = 8
	preCounted[19] = 8
	preCounted[20] = 6
	preCounted[30] = 6
	preCounted[40] = 5
	preCounted[50] = 5
	preCounted[60] = 5
	preCounted[70] = 7
	preCounted[80] = 6
	preCounted[90] = 6
	sum := 0
	fmt.Println(countLetters(342))
	fmt.Println(countLetters(115))
	fmt.Println(countLetters(12))
	fmt.Println(countLetters(81))
	fmt.Println(countLetters(49))
	for i := 1; i <= 999; i++ {
		sum += countLetters(i)
	}
	fmt.Println(sum + 11)
}
