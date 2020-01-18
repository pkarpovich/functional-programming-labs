package main

import (
	"fmt"
	"functional-programming-labs/lab-1/tasks"
)

func main() {
	const row = 2
	const cell = 1
	number := tasks.GetNumberFromPascaleTriangle(row, cell)

	fmt.Printf("Pascale(%d,%d): %d", row, cell, number)
}