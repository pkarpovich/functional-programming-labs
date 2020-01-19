package main

import (
	"fmt"
	"functional-programming-labs/lab-1/tasks"
)

func main() {
	const row = 2
	const cell = 1
	number := tasks.GetNumberFromPascaleTriangle(row, cell)

	fmt.Printf("Pascale(%d,%d): %d\n", row, cell, number)

	coins := []int{1, 2, 3, 4, 5}
	combinationCount := tasks.GetAllCombinations(10, len(coins)-1, coins)

	fmt.Printf("Count of combinations: %d", combinationCount)

}
