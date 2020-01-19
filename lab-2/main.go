package main

import (
	"fmt"
	"functional-programming-labs/lab-2/tasks"
)

func main() {
	sequence := []int{1, 2, 3, 4, 1, 2, 3, 4, 5, 1, 2, 3, 1, 2, 3, 4, 5, 6}

	maxIncreasingSubsequence := tasks.FindMaxIncreasingSubsequenceImperative(sequence)

	fmt.Printf("Imperative\n\n"+
		"Original sequence: %v\n"+
		"Max Increasing Subsequence: %v\n",
		sequence, maxIncreasingSubsequence)

	maxIncreasingSubsequence = tasks.FindMaxIncreasingSubsequenceFunctional(sequence)

	fmt.Printf("Functional\n\n"+
		"Original sequence: %v\n"+
		"Max Increasing Subsequence: %v\n",
		sequence, maxIncreasingSubsequence)
}
