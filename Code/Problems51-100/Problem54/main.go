package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	value int
	suit  byte
}

func sortCardsByValue(cards []card) []card {
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			if cards[j].value < cards[j-1].value {
				cards[j], cards[j-1] = cards[j-1], cards[j]
			}
		}
	}
	return cards
}

func sortCardsBySuit(cards []card) []card {
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			if cards[j].suit < cards[j-1].suit {
				cards[j], cards[j-1] = cards[j-1], cards[j]
			}
		}
	}
	return cards
}

func disposeCards(hand string, p1, p2 *[]card) {
	cardsStrings := strings.Split(hand, " ")
	for i := 0; i < 5; i++ {
		*p1 = append(*p1, decodeCard(cardsStrings[i]))
		*p2 = append(*p2, decodeCard(cardsStrings[5+i]))
	}
}
func decodeCard(str string) card {
	var crd card
	realValues := make(map[string]int)
	realValues["J"] = 11
	realValues["Q"] = 12
	realValues["K"] = 13
	realValues["A"] = 14
	realValues["T"] = 10

	crd.suit = str[len(str)-1]

	valByte := str[:1]
	if valByte == "T" || valByte == "J" || valByte == "A" || valByte == "Q" || valByte == "K" {
		crd.value = realValues[valByte]
		return crd
	}
	val, err := strconv.Atoi(valByte)
	if err != nil {
		panic(err)
	}
	crd.value = val

	return crd
}

func getBestHand(cards []card) (int, []int) {
	if indexes := hasStraightFlush(cards); indexes != nil {
		if flushIsRoyal(cards) {
			return 10, indexes
		}
		return 9, indexes
	}
	if indexes := hasFourOfAKind(cards); indexes != nil {
		return 8, indexes
	}
	if indexes := hasFullHouse(cards); indexes != nil {
		return 7, indexes
	}
	if indexes := hasFlush(cards); indexes != nil {
		return 6, indexes
	}
	if indexes := hasStraight(cards); indexes != nil {
		return 5, indexes
	}
	if indexes := hasThree(cards); indexes != nil {
		return 4, indexes
	}
	if indexes := hasTwoPairs(cards); indexes != nil {
		return 3, indexes
	}
	if indexes := hasPair(cards); indexes != nil {
		return 2, indexes
	}
	return 1, getBestCard(cards)
}

func countEachValue(cards []card) []int {
	vals := make([]int, 15)
	for i := range cards {
		vals[cards[i].value]++
	}
	return vals
}

func compareValues(p1 []card, p2 []card) int {
	for i := range p1 {
		if p1[i].value > p2[i].value {
			return 1
		}
		if p2[i].value > p1[i].value {
			return 2
		}
	}
	return 0
}

func compareHouses(p1, p2 []card) int {
	p1ct := countEachValue(p1)
	p2ct := countEachValue(p2)
	var valp1, valp2 int

	for i := 1; i <= 14; i++ {
		if p1ct[i] == 3 {
			valp1 = i
		}
		if p2ct[i] == 3 {
			valp2 = i
		}
	}
	if valp1 > valp2 {
		return 1
	}
	if valp2 > valp1 {
		return 2
	}

	for i := 1; i <= 14; i++ {
		if p1ct[i] == 2 {
			valp1 = i
		}
		if p2ct[i] == 2 {
			valp2 = i
		}
	}
	if valp1 > valp2 {
		return 1
	}
	if valp2 > valp1 {
		return 2
	}

	return 0
}

func getWinner(p1, p2 []card) int {
	bestP1, p1Indexes := getBestHand(p1)
	bestP2, p2Indexes := getBestHand(p2)
	if bestP1 != bestP2 {
		if bestP1 > bestP2 {
			return 1
		}
		return 2
	}
	var comparedHands int
	if bestP1 == 7 {
		comparedHands = compareHouses(p1, p2)
	} else {
		comparedHands = compareValues(pickCards(p1, p1Indexes), pickCards(p2, p2Indexes))
	}
	if comparedHands != 0 {
		return comparedHands
	}
	return compareValues(removeCards(p1, p1Indexes), removeCards(p2, p2Indexes))
}

func pickCards(cards []card, indexes []int) []card {
	var newCards []card
	for _, i := range indexes {
		newCards = append(newCards, cards[i])
	}
	return newCards

}

func removeCards(cards []card, indexes []int) []card {
	cardsRemoved := 0
	sort.Ints(indexes)
	for i := range indexes {
		cards = append(cards[:i+cardsRemoved], cards[i+cardsRemoved+1:]...)
		cardsRemoved++
	}
	return cards
}

func main() {
	wins := 0
	file, _ := os.Open("p054_poker.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var p1, p2 []card
		disposeCards(scanner.Text(), &p1, &p2)
		if getWinner(p1, p2) == 1 {
			wins++
		}
	}
	fmt.Println(wins)
}
