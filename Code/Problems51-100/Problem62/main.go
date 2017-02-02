package main

import (
	"fmt"
	"math/big"
	"sort"
)

type sortRunes []rune
type pair struct {
	counted  int
	smallest string
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	m := make(map[string]pair)
	c := big.NewInt(1)
	for i := 1; ; i++ {
		s := SortString(c.String())
		if _, ok := m[s]; !ok {
			m[s] = pair{1, c.String()}
		} else {
			m[s] = pair{m[s].counted + 1, m[s].smallest}
		}
		if m[s].counted == 5 {
			fmt.Println(m[s].smallest)
			return
		}
		j := int64(i)
		add := big.NewInt(j*j*3 + j*3 + 1)
		c.Add(c, add)
	}
}
