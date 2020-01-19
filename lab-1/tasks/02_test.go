package tasks

import (
"github.com/stretchr/testify/assert"
"testing"
)

func TestGetAllCombinations(t *testing.T) {
	tests := []struct {
		money  int
		coins  []int
		answer int
	}{
		{4, []int{1, 2}, 3},
		{4, []int{1, 2, 3}, 4},
		{10, []int{1, 2, 3, 4, 5}, 30},
	}

	for _, test := range tests {
		assert.Equal(t, GetAllCombinations(test.money, len(test.coins)-1, test.coins), test.answer)
	}
}
