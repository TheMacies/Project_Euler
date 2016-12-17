package main

import "fmt"

func main() {
	currentDay := 1
	sundays := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			if currentDay == 6 {
				sundays++
			}
			if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
				currentDay += 31
				currentDay %= 7
				continue
			}
			if month == 2 {
				if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
					currentDay += 29
					currentDay %= 7
					continue
				}
				currentDay += 28
				currentDay %= 7
				continue
			}
			currentDay += 30
			currentDay %= 7
		}
	}
	fmt.Println(sundays)
}
