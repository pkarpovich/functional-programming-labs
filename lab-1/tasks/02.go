package tasks

func GetAllCombinations(money int, currentCoinIndex int, coins []int) int {
	if money < 0 || currentCoinIndex < 0 {
		return 0
	}
	if money == 0 || currentCoinIndex == 0 {
		return 1
	}

	return GetAllCombinations(money, currentCoinIndex-1, coins) +
		GetAllCombinations(money-coins[currentCoinIndex], currentCoinIndex, coins)
}
