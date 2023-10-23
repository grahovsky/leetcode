package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test66(t *testing.T) {
	tests := []struct {
		input  []int
		result []int
	}{
		{input: []int{1, 2, 3}, result: []int{1, 2, 4}},
		{input: []int{4, 3, 2, 1}, result: []int{4, 3, 2, 2}},
		{input: []int{9}, result: []int{1, 0}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 58", func(t *testing.T) {
			result := plusOne(tc.input)
			require.Equal(t, tc.result, result)
		})
	}
}
