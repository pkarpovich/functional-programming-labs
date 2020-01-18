package tasks

func GetNumberFromPascaleTriangle(row int, cell int) int {
	return factorial(row) / (factorial(cell) * factorial(row-cell))
}

func factorial(n int) (result int) {
	if n > 0 {
		return n * factorial(n-1)
	}

	return 1
}