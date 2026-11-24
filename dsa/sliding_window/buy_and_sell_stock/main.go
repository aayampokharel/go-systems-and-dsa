package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	selectedLeastVal := prices[0]
	maxProfit, tempResult := 0, 0
	for i := 1; i < len(prices); i++ {
		if selectedLeastVal > prices[i] {
			selectedLeastVal = prices[i]
		} else if selectedLeastVal < prices[i] {
			tempResult = prices[i] - selectedLeastVal
			if tempResult > maxProfit {
				maxProfit = tempResult
			}
		}
	}
	return maxProfit
}

func BetterSolution(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
