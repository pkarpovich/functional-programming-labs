package tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxIncreasingSubsequenceImperative(t *testing.T) {
	tests := []struct {
		sequence []int
		answer   []int
	}{
		{[]int{1, 2, 3, 4, 1, 2, 3, 4, 5, 1, 2, 3, 1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		assert.Equal(t, FindMaxIncreasingSubsequenceImperative(test.sequence), test.answer)
	}
}

func TestFindMaxIncreasingSubsequenceFunctional(t *testing.T) {
	tests := []struct {
		sequence []int
		answer   []int
	}{
		{[]int{1, 2, 3, 4, 1, 2, 3, 4, 5, 1, 2, 3, 1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		assert.Equal(t, FindMaxIncreasingSubsequenceFunctional(test.sequence), test.answer)
	}
}
