package main

func getBestCard(cards []card) []int {
	var ind []int
	ct := countEachValue(cards)
	biggestVal := 1
	for i := 14; i >= 1; i-- {
		if ct[i] != 0 {
			biggestVal = i
			break
		}
	}
	for i := range cards {
		if cards[i].value == biggestVal {
			ind = append(ind, i)
			return ind
		}
	}
	return nil
}

func hasPair(cards []card) []int {
	var ind []int
	ct := countEachValue(cards)
	pairVal := -1
	for i := 14; i >= 1; i-- {
		if ct[i] == 2 {
			pairVal = i
			break
		}
	}

	if pairVal == -1 {
		return nil
	}
	for i := range cards {
		if cards[i].value == pairVal {
			ind = append(ind, i)
		}
	}
	return ind
}

func hasThree(cards []card) []int {
	var ind []int
	ct := countEachValue(cards)
	pairVal := -1
	for i := 14; i >= 1; i-- {
		if ct[i] == 3 {
			pairVal = i
			break
		}
	}
	if pairVal == -1 {
		return nil
	}
	for i := range cards {
		if cards[i].value == pairVal {
			ind = append(ind, i)
		}
	}
	return ind
}

func hasTwoPairs(cards []card) []int {
	var ind []int
	ct := countEachValue(cards)
	pair1Val := -1
	pair2Val := -1
	for i := 14; i >= 1; i-- {
		if ct[i] == 2 {
			if pair1Val == -1 {
				pair1Val = i
			} else {
				pair2Val = i
				break
			}
		}
	}
	if pair2Val == -1 {
		return nil
	}
	for i := range cards {
		if cards[i].value == pair1Val || cards[i].value == pair2Val {
			ind = append(ind, i)
		}
	}
	return ind
}

func hasStraight(cards []card) []int {
	var ind []int
	ct := countEachValue(cards)
	for i := 0; i <= 14; i++ {
		if ct[i] != 0 {
			for j := 1; j < 5; j++ {
				if ct[i+j] != 1 {
					return nil
				}
			}
			break
		}
	}
	for i := range cards {
		ind = append(ind, i)
	}
	return ind
}

func hasFlush(cards []card) []int {
	var ind []int
	for i := 0; i < 4; i++ {
		if cards[i].suit != cards[i+1].suit {
			return nil
		}
	}
	for i := range cards {
		ind = append(ind, i)
	}
	return ind
}

func hasFullHouse(cards []card) []int {
	three := hasThree(cards)
	if three == nil {
		return nil
	}
	pair := hasPair(cards)
	if pair == nil {
		return nil
	}
	return append(three, pair...)
}

func hasFourOfAKind(cards []card) []int {
	var ind []int
	ct := countEachValue(cards)
	pairVal := 1
	for i := 14; i >= 1; i-- {
		if ct[i] == 4 {
			pairVal = i
			break
		}
	}
	for i := range cards {
		if cards[i].value == pairVal {
			ind = append(ind, i)
		}
	}
	return ind
}

func hasStraightFlush(cards []card) []int {
	flush := hasFlush(cards)
	if flush == nil {
		return nil
	}
	straight := hasStraight(cards)
	if straight == nil {
		return nil
	}
	return flush
}

func flushIsRoyal(cards []card) bool {
	for i := range cards {
		if cards[i].value == 14 {
			return true
		}
	}
	return false
}
