package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test35(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		result int
	}{
		{input: []int{1, 3, 5, 6}, target: 5, result: 2},
		{input: []int{1, 3, 5, 6}, target: 2, result: 1},
	}
	for _, tc := range tests {
		tc := tc
		t.Run("test 28", func(t *testing.T) {
			result := searchInsert(tc.input, tc.target)
			require.Equal(t, tc.result, result)
		})
	}
}
