package main

import (
	"fmt"
)

const BlockLen = 50

func calculateCombinations(colorLen int64) int64 {
	maxUsing := make([]int64,BlockLen+1)
	for i:= colorLen;i<=BlockLen;i++ {
	 maxUsing[i] = 1 + maxUsing[i-1] + maxUsing[i-colorLen]
	}
	return maxUsing[BlockLen]
}

func main() {
	fmt.Println(calculateCombinations(2)  +calculateCombinations(3)  +calculateCombinations(4))
}

