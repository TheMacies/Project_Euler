package main

import "fmt"

func sumDiagonals(spiral [][]int) int {
	sum := 0
	for i := 0; i < len(spiral); i++ {
		sum += spiral[i][len(spiral)-i-1]
		sum += spiral[i][i]
	}
	// substraction because element in the middle was added twice
	return sum - 1
}

func generateSpiral(side int) [][]int {
	// cannot make a spiral out of even-length side square
	if side%2 == 0 {
		side++
	}

	spiral := make([][]int, side)
	for i := 0; i < side; i++ {
		spiral[i] = make([]int, side)
	}
	elemets := side * side

	x := side / 2
	y := side / 2
	spiral[x][y] = 1
	currentValue := 2
	currentSide := 3
	for {
		if elemets < currentValue {
			break
		}
		x++
		for i := 0; i < currentSide-2; i++ {
			spiral[x][y] = currentValue
			currentValue++
			y--
		}
		for i := 0; i < currentSide-1; i++ {
			spiral[x][y] = currentValue
			currentValue++
			x--
		}
		for i := 0; i < currentSide-1; i++ {
			spiral[x][y] = currentValue
			currentValue++
			y++
		}
		for i := 0; i < currentSide-1; i++ {
			spiral[x][y] = currentValue
			currentValue++
			x++
		}

		spiral[x][y] = currentValue
		currentValue++

		currentSide += 2
	}
	return spiral
}

func main() {
	spiral := generateSpiral(1001)
	fmt.Println(sumDiagonals(spiral))
}
