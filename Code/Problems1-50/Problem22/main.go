package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func evaluate(name string) int {
	sum := 0
	for _, letter := range name {
		sum += (int(letter) - 64)
	}
	return sum
}

func main() {
	content, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	sum := 0
	namesLine := string(content)
	names := strings.Split(namesLine, ",")
	sort.Strings(names)
	for i, name := range names {
		ev := evaluate(name[1 : len(name)-1])
		fmt.Println(i+1, ev, name)
		sum += (i + 1) * ev
	}
	fmt.Println(sum)
}
