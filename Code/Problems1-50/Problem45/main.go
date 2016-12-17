package main

import "fmt"

func triang(n int) int {
	return (n * (n + 1)) / 2
}

func penta(n int) int {
	return (n * (3*n - 1)) / 2
}

func hexa(n int) int {
	return (n * (2*n - 1))
}

func main() {
	// Init values taken from Euler page
	tInd := 286
	tVal := triang(tInd)
	pInd := 166
	pVal := penta(pInd)
	hInd := 144
	hVal := hexa(hInd)
	for {
		for tVal < pVal || tVal < hVal {
			tInd++
			tVal = triang(tInd)
		}

		for pVal < tVal || pVal < hVal {
			pInd++
			pVal = penta(pInd)
		}

		for hVal < pVal || hVal < tVal {
			hInd++
			hVal = hexa(hInd)
		}
		if pVal == hVal && hVal == tVal {
			break
		}
	}
	fmt.Println(hVal)
}
