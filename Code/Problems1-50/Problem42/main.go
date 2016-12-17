package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func evaluateWord(word string) int {
	val := 0
	for c := range word {
		val += int(word[c]) - 64
	}
	return val
}

func main() {
	var triangularsCount int
	var triangulars []int
	for i := 1; i < 100; i++ {
		triangulars = append(triangulars, i*(i+1)/2)
	}

	file, err := ioutil.ReadFile("p042_words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	words := strings.Split(string(file), ",")
	fmt.Println(triangulars)
	for i := range words {
		eval := evaluateWord(words[i][1 : len(words[i])-1])

		for j := range triangulars {
			if triangulars[j] == eval {
				triangularsCount++
				break
			}
		}
	}
	fmt.Println(triangularsCount)
}
