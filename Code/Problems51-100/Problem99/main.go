package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type number struct {
	base, exp int
}

func pow(base, exp int) *big.Int {
	res := big.NewInt(1)
	return res.Exp(big.NewInt(int64(base)), big.NewInt(int64(exp)), nil)
}

func compare(n1, n2 number) int {
	a := math.Log(float64(n1.base)) * float64(n1.exp)
	b := math.Log(float64(n2.base)) * float64(n2.exp)
	if a < b {
		return -1
	}
	return 1
}

func findGreatest(nums []number) number {
	max := nums[0]
	for i := range nums[1:] {
		fmt.Println(i)
		if compare(max, nums[i]) == -1 {
			max = nums[i]
		}
	}
	return max
}

func main() {
	f, _ := os.Open("p099_base_exp.txt")
	bytes, _ := ioutil.ReadAll(f)
	rows := strings.Split(string(bytes), "\n")
	var nums []number
	for _, str := range rows {
		if str == "" {
			break
		}
		separated := strings.Split(str, ",")
		base, _ := strconv.Atoi(separated[0])
		exp, _ := strconv.Atoi(separated[1])
		nums = append(nums, number{base, exp})
	}

	fmt.Println(findGreatest(nums))
}
