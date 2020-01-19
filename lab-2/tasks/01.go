package tasks

func FindMaxIncreasingSubsequenceImperative(sequence []int) (maxSubsequence []int) {
	var prevItem int

	var currentSubsequence []int

	for index, item := range sequence {
		if prevItem < item && index != len(sequence)-1 {
			prevItem = item

			currentSubsequence = append(currentSubsequence, item)
			continue
		}
		if index == len(sequence)-1 {
			currentSubsequence = append(currentSubsequence, item)
		}

		if len(currentSubsequence) > len(maxSubsequence) {
			maxSubsequence = append([]int{}, currentSubsequence...)
		}
	}

	return
}

func FindMaxIncreasingSubsequenceFunctional(sequence []int) []int {
	for index, item := range sequence {
		if index == 0 || sequence[index-1] < item {
			continue
		}

		currentSubsequence := sequence[:index]
		nextSubsequence := FindMaxIncreasingSubsequenceFunctional(sequence[index:])

		if len(currentSubsequence) > len(nextSubsequence) {
			return currentSubsequence
		}

		return nextSubsequence
	}

	return sequence
}
