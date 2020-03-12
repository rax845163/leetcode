package q122

func maxProfit(prices []int) int {
	return toBuy(0, prices, make(map[int]int))
}

func toBuy(day int, prices []int, maxProfitMap map[int]int) int {
	if day >= len(prices)-1 {
		return 0
	}
	if p, exist := maxProfitMap[day]; exist {
		return p
	}

	buy := toSell(day+1, prices, prices[day], maxProfitMap)
	notBuy := toBuy(day+1, prices, maxProfitMap)

	max := buy
	if notBuy > buy {
		max = notBuy
	}
	maxProfitMap[day] = max
	return max
}

func toSell(day int, prices []int, buyIn int, maxProfitMap map[int]int) int {
	if day == len(prices) {
		return 0
	}
	if p, exist := maxProfitMap[day]; exist {
		return p
	}

	sell := (prices[day] - buyIn) + toBuy(day+1, prices, maxProfitMap)
	keep := toSell(day+1, prices, buyIn, maxProfitMap)

	max := sell
	if keep > sell {
		max = keep
	}
	maxProfitMap[day] = max
	return max
}
