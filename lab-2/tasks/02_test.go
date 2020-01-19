package tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchImperative(t *testing.T) {
	tests := []struct {
		array  []int
		target int
		answer int
	}{
		{[]int{1, 2, 4, 5, 7, 8, 9}, 9, 6},
		{[]int{1, 2, 4, 5, 7, 8, 9}, 1, 0},
		{[]int{1, 2, 4, 5, 7, 8, 9}, 4, 2},
	}

	for _, test := range tests {
		assert.Equal(t, BinarySearchImperative(test.array, test.target), test.answer)
	}
}

func TestBinarySearchFunctional(t *testing.T) {
	tests := []struct {
		array  []int
		target int
		answer int
	}{
		{[]int{1, 2, 4, 5, 7, 8, 9}, 9, 6},
		{[]int{1, 2, 4, 5, 7, 8, 9}, 1, 0},
		{[]int{1, 2, 4, 5, 7, 8, 9}, 4, 2},
	}

	for _, test := range tests {
		assert.Equal(t, BinarySearchFunctional(test.array, test.target, 0, len(test.array)), test.answer)
	}
}
