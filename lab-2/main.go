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
		"Max Increasing Subsequence: %v\n\n",
		sequence, maxIncreasingSubsequence)

	arr := []int{1, 2, 4, 5, 7, 8, 9}
	target := 9

	targetIndex := tasks.BinarySearchImperative(arr, target)

	fmt.Printf("Imperative\n\n"+
		"Array: %v\n"+
		"Target: %d\n"+
		"Target index: %d\n", arr, target, targetIndex)

	targetIndex = tasks.BinarySearchFunctional(arr, target, 0, len(arr))

	fmt.Printf("Imperative\n\n"+
		"Array: %v\n"+
		"Target: %d\n"+
		"Target index: %d\n", arr, target, targetIndex)
}
