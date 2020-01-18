package tasks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNumberFromPascaleTriangle(t *testing.T) {
	tests := []struct {
		row    int
		cell   int
		answer int
	}{
		{2, 0, 1},
		{2, 1, 2},
		{3, 1, 3},
	}

	for _, test := range tests {
		assert.Equal(t, GetNumberFromPascaleTriangle(test.row, test.cell), test.answer)
	}
}