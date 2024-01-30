package main

import (
	"fmt"
)

func birthdayCandles(candles []int) int {
	maxHeight := candles[0] // Initialize maxHeight with the height of the first candle

	// Loop through the candles to find the tallest candle height
	for _, height := range candles {
		if height > maxHeight {
			maxHeight = height
		}
	}

	count := 0
	// Count how many candles have the tallest height
	for _, height := range candles {
		if height == maxHeight {
			count++
		}
	}

	return maxHeight
}

func countMax(candles []int, param int) int { // count total max height

	max := candles[0] // Initialize maxHeight with the height of the first candle

	// Iterate through the array
	for _, num := range candles {
		// If current number is greater than max, update max
		if num > max {
			max = num
		}
	}

	// return max
	count := 0
	for _, num := range candles {
		if num == param {
			count++
		}
	}
	return count
}

func main() {
	candles := []int{4, 4, 1, 3, 4}
	fmt.Println("list: ", candles)
	fmt.Println("Number of candles that can be blown out:", birthdayCandles(candles))
	fmt.Println("max height value: ", countMax(candles, birthdayCandles(candles)))
}
