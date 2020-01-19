package tasks

func BinarySearchImperative(array []int, target int) int {
	first := 0
	last := len(array) - 1

	for first <= last {
		mid := (first + last) / 2

		if array[mid] == target {
			return mid
		}

		if array[mid] < target {
			first = mid + 1
		}

		if array[mid] > target {
			last = mid - 1
		}
	}

	return -1
}

func BinarySearchFunctional(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex

	var mid int

	for startIndex < endIndex {
		mid = (lowIndex + highIndex) / 2
		if array[mid] > target {
			return BinarySearchFunctional(array, target, lowIndex, mid)
		} else if array[mid] < target {
			return BinarySearchFunctional(array, target, mid+1, highIndex)
		} else {
			return mid
		}

	}
	return -1

}
